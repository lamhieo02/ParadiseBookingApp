package bookingguiderhandler

import (
	"net/http"
	"paradise-booking/constant"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) GetBookingGuiderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, _ := c.GetQuery("booking_guider_id")

		bookingId, _ := strconv.Atoi(bookingID)

		err := hdl.bookingGuiderUC.GetBookingByID(c, bookingId)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingFail)
			return
		}

		c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingSuccess)
	}
}
