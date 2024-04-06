package postreviewhandler

import (
	"net/http"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) UpdatePostReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReviewBody postreviewiomodel.UpdatePostReviewReq

		if err := c.ShouldBind(&postReviewBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}

		err := hdl.postReviewUC.UpdatePostReview(c.Request.Context(), &postReviewBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
