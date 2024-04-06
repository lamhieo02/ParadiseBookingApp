package postreviewhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) DeletePostReviewByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		postReviewID := c.Param("post_review_id")
		id, err := strconv.Atoi(postReviewID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.postReviewUC.DeletePostReviewByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
