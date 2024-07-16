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

func (hdl *apiUtilHandler) AggregateDataPostReview() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		postReviewIDS := ctx.Query("post_review_ids")

		arrPostReviewID := strings.Split(postReviewIDS, ",")

		var result []string

		for _, postReviewID := range arrPostReviewID {
			id, err := strconv.Atoi(postReviewID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			postReview, err := hdl.postReviewSto.GetByID(ctx, id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			if postReview == nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "postReview not found"})
				return
			}

			// aggregate data
			res := aggregatePostReview(postReview)
			result = append(result, res)
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func aggregatePostReview(postReview *entities.PostReview) string {
	res := ""
	res += fmt.Sprintf("Chủ bài viết có id: %v, bài review có title là: %s, chủ đề của bài review là: %s, có content nội dung là: %s, địa điểm ở %s - %s - %s",
		postReview.PostOwnerId, postReview.Title, constant.MapCategoryIDToName[postReview.Topic], postReview.Content, postReview.Country, postReview.State, postReview.District)

	return res
}
