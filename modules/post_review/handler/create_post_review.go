package postreviewhandler

import (
	"net/http"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) CreatePostReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReviewBody postreviewiomodel.CreatePostReviewReq

		if err := c.ShouldBind(&postReviewBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.postReviewUC.CreatePostReview(c.Request.Context(), &postReviewBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
