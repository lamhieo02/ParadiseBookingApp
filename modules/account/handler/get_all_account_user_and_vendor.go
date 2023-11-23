package accounthandler

import (
	"net/http"
	"paradise-booking/modules/account/convert"

	"github.com/gin-gonic/gin"
)

func (hdl *accountHandler) GetAllAccountUserAndVendor() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := hdl.accountUC.GetAllAccountUserAndVendor(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res := convert.ConvertAccountEntityToInfoMangageForAdmin(result)
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
