package apiutils

import (
	"fmt"
	"net/http"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (hdl *apiUtilHandler) AggregateDataPostGuide() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postGuideIDS := ctx.Query("post_guide_ids")

		arrPostGuideID := strings.Split(postGuideIDS, ",")

		var result []string

		for _, postGuideID := range arrPostGuideID {
			id, err := strconv.Atoi(postGuideID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			postGuide, err := hdl.postGuideCache.GetByID(ctx.Request.Context(), id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			if postGuide == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "postGuide not found"})
				return
			}

			// get place related to post guide
			condition := map[string]interface{}{
				"state": postGuide.State,
			}
			placeIds, err := hdl.placeSto.ListPlaceIdsByCondition(ctx, 10, condition)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			var places []*entities.Place
			for _, placeID := range placeIds {
				place, err := hdl.placeCache.GetPlaceByID(ctx, placeID)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
					return
				}
				places = append(places, place)
			}

			// aggregate data
			res := aggregatePostGuide(postGuide, places)
			result = append(result, res)
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func aggregatePostGuide(postGuide *entities.PostGuide, places []*entities.Place) string {
	res := ""
	// write template like below
	postGuideURL := "https://booking.workon.space/post-guiders/" + strconv.Itoa(postGuide.Id)
	res = fmt.Sprintf("Bài viết có tiêu đề %s về chủ đề %s, có mô tả: %s, có thông tin chi tiết như sau: \n Quốc gia: %s, Tỉnh/Thành phố: %s, Quận/Huyện: %s, Địa chỉ: %s, Lịch trình: %s, Ngôn ngữ: %s, Ngày tạo: %s, Ngày cập nhật: %s, thông tin của post_guide này khi được hỏi nên được trả về là đường dẫn tới post_guide đó như sau: %s",
		postGuide.Title, constant.MapPostGuideTopic[postGuide.TopicID], postGuide.Description, postGuide.Country, postGuide.State, postGuide.District, postGuide.Address, postGuide.Schedule, postGuide.Languages, postGuide.CreatedAt, postGuide.UpdatedAt, postGuideURL)
	// get post guide related to place
	res += fmt.Sprintln()
	res += fmt.Sprintf("Các Places liên quan tới PostGuide này, liên quan ở đây tức là place/hotel/homestay đó có vị trí gần với post_guide này diễn ra, thông tin của các place(địa điểm/hotel/homestay) đó theo đường dẫn trực tiếp tới place đó như sau:")
	for _, place := range places {
		placeURL := "https://booking.workon.space/listings/" + strconv.Itoa(place.Id)
		res += fmt.Sprintf("%s", placeURL)
		res += fmt.Sprintln(" ")
	}

	return res
}
