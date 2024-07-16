package apiutils

import (
	"errors"
	"fmt"
	"net/http"
	"paradise-booking/entities"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (hdl *apiUtilHandler) AggregateDataPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeIDS := ctx.Query("place_ids")

		arrPlaceID := strings.Split(placeIDS, ",")

		var result []string

		for _, placeID := range arrPlaceID {
			id, err := strconv.Atoi(placeID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			place, err := hdl.placeSto.GetPlaceByID(ctx.Request.Context(), id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			if place == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Place not found"})
				return
			}

			// get post guide related to place
			condition := map[string]interface{}{
				"state": place.State,
			}
			postGuideIds, err := hdl.postGuideSto.ListPostGuideIdsByCondition(ctx, 10, condition)
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			var postGuides []*entities.PostGuide
			for _, postGuideID := range postGuideIds {
				postGuide, err := hdl.postGuideCache.GetByID(ctx, postGuideID)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
					return
				}
				postGuides = append(postGuides, postGuide)
			}

			// aggregate data
			res := aggregatePlace(place, postGuides)
			result = append(result, res)
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func aggregatePlace(place *entities.Place, postGuide []*entities.PostGuide) string {
	res := ""
	// write template like below
	placeURL := "https://booking.workon.space/listings/" + strconv.Itoa(place.Id)
	res = fmt.Sprintf("Home %s là một homestay theo mô tả như sau: %s, nằm tại địa chỉ %s, hotel/home này có giá thuê là %v một đêm và có thể đón tiếp tối đa %v khách. Với không gian %v phòng ngủ và %v giường. Địa chỉ: %s, Tọa độ: vĩ độ:%v, kinh độ: %v, thông tin bổ sung: giá phòng 1 đêm: %v, mã định dạng / mã id của địa điểm này: %v, thông tin của place này khi được hỏi nên được trả về là đường dẫn tới place đó như sau: %s",
		place.Name, place.Description, place.Address, place.PricePerNight, place.MaxGuest, place.BedRoom, place.NumBed, place.Address, place.Lat, place.Lng, place.PricePerNight, place.Id, placeURL)

	res += fmt.Sprintln()
	// get post guide related to place
	for _, postGuide := range postGuide {
		res += fmt.Sprintf("Các Post Guide liên quan tới Place này, liên quan ở đây tức là tour du lịch đó có vị trí diễn ra gần với place này, thông tin của các post_guide(bài đăng hướng dẫn viên) đó theo đường dẫn trực tiếp tới bài post đó như sau:")
		postGuideURL := "https://booking.workon.space/post-guiders/" + strconv.Itoa(postGuide.Id)
		res += fmt.Sprintf("%s", postGuideURL)
		res += fmt.Sprintln(" ")
	}

	return res
}
