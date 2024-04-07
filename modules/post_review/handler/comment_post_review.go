package postreviewhandler

import (
	"net/http"
	postreviewratingiomodel "paradise-booking/modules/post_review_rating/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) CommentPostReview() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postReviewRatingBody postreviewratingiomodel.CommentPostReviewRatingReq
		if err := c.ShouldBind(&postReviewRatingBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}

		err := hdl.postReviewUC.CommentPostReview(c.Request.Context(), &postReviewRatingBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
