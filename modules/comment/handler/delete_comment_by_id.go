package commenthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *commentHandler) DeleteCommentByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		id, err := strconv.Atoi(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.commentUC.DeleteCommentByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
