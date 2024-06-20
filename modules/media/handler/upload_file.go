package mediahandler

import (
	"net/http"
	"paradise-booking/common"

	"github.com/gin-gonic/gin"
)

func (hdl *mediaHandler) UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		files := form.File["files"]

		img, err := hdl.mediaUC.UploadFile(c.Request.Context(), files)
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		c.JSON(http.StatusOK, gin.H{"data": img})
	}
}
