package bookingguiderhandler

import (
	"net/http"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) CreateBookingGuider() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bookingguideriomodel.CreateBookingReq
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingGuiderUC.CreateBookingGuider(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
