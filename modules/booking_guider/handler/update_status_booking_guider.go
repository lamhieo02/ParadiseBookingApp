package bookingguiderhandler

import (
	"net/http"
	"paradise-booking/constant"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) ConfirmStatusBookingGuider() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, _ := c.GetQuery("booking_guider_id")
		status, _ := c.GetQuery("status")

		bookingId, _ := strconv.Atoi(bookingID)
		statusInt, _ := strconv.Atoi(status)

		err := hdl.bookingGuiderUC.UpdateStatusBooking(c.Request.Context(), bookingId, statusInt)
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingFail)
			return
		}

		c.Redirect(http.StatusMovedPermanently, constant.UrlConfirmBookingSuccess)
	}
}

func (hdl *bookingGuiderHandler) UpdateStatusBookingGuider() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, _ := c.GetQuery("booking_guider_id")
		status, _ := c.GetQuery("status")

		bookingId, _ := strconv.Atoi(bookingID)
		statusInt, _ := strconv.Atoi(status)

		err := hdl.bookingGuiderUC.UpdateStatusBooking(c.Request.Context(), bookingId, statusInt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
