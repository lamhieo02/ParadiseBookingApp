package replycommenthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *replyCommentHandler) DeleteReplyComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		replyCommentID := c.Param("reply_comment_id")

		id, err := strconv.Atoi(replyCommentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.replyCommentUC.DeleteByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
