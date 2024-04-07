package replycommenthandler

import (
	"net/http"
	replycommentiomodel "paradise-booking/modules/reply_comment/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *replyCommentHandler) ReplySourceComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var replySourceComment replycommentiomodel.ReplyCommentReq
		if err := c.ShouldBind(&replySourceComment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.replyCommentUC.ReplySourceComment(c.Request.Context(), &replySourceComment)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
