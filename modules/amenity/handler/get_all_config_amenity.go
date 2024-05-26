package amenityhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) GetAllConfigAmenity() gin.HandlerFunc {
	return func(c *gin.Context) {
		_type := c.Query("type")
		typeInt, _ := strconv.Atoi(_type)
		res, err := hdl.amenityUC.GetAllConfigAmenity(c.Request.Context(), typeInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})

	}
}
