package mediahandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hdl *mediaHandler) GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Param("path")

		c.Header("Content-Type", "image/png")
		http.ServeFile(c.Writer, c.Request, "images/"+path)
	}
}
