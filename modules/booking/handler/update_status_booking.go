package bookinghandler

import (
	"net/http"
	"paradise-booking/modules/booking/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) UpdateStatusBooking() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data iomodel.UpdateStatusBookingReq
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		err := hdl.bookingUC.UpdateStatusBooking(c.Request.Context(), data.BookingID, data.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
