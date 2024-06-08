package commenthandler

import (
	"net/http"
	commentiomodel "paradise-booking/modules/comment/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *commentHandler) EditCommentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		id, err := strconv.Atoi(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var updateBody commentiomodel.UpdateCommentReq
		if err := c.ShouldBind(&updateBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.commentUC.EditCommentByID(c.Request.Context(), id, updateBody.ToEntity())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
