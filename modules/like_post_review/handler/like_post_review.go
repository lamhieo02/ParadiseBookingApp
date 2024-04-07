package likepostreviewhandler

import (
	"net/http"
	likepostreviewiomodel "paradise-booking/modules/like_post_review/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *likePostReviewHandler) LikePostReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var likePostReviewBody likepostreviewiomodel.LikePostReviewReq
		if err := c.ShouldBind(&likePostReviewBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}

		err := hdl.likePostReview.LikePostReview(c.Request.Context(), &likePostReviewBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
