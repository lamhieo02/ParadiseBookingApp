package postreviewhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *postReviewHandler) GetPostReviewByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		postReviewID := c.Param("post_review_id")
		id, err := strconv.Atoi(postReviewID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := hdl.postReviewUC.GetPostReviewByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
