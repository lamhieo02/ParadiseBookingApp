package commenthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *commentHandler) GetCommentByPostReviewID() gin.HandlerFunc {
	return func(c *gin.Context) {
		PostReviewID := c.Param("post_review_id")

		id, err := strconv.Atoi(PostReviewID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := hdl.commentUC.GetCommentByPostReviewID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
