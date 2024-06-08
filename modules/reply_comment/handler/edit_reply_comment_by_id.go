package replycommenthandler

import (
	"net/http"
	replycommentiomodel "paradise-booking/modules/reply_comment/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *replyCommentHandler) EditReplyCommentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		replyCommentID := c.Param("reply_comment_id")

		id, err := strconv.Atoi(replyCommentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateBody replycommentiomodel.UpdateReplyCommentReq
		if err := c.ShouldBind(&updateBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.replyCommentUC.EditReplyCommentByID(c.Request.Context(), id, updateBody.ToEntity())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
