package mediahandler

import (
	"net/http"
	"paradise-booking/common"

	"github.com/gin-gonic/gin"
)

func (hdl *mediaHandler) UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		img, err := hdl.mediaUC.UploadFile(c.Request.Context(), fileHeader)
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		c.JSON(http.StatusOK, gin.H{"data": img})
	}
}

func (hdl *mediaHandler) GetImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Param("path")

		c.Header("Content-Type", "image/png")
		http.ServeFile(c.Writer, c.Request, "images/"+path)
	}
}
