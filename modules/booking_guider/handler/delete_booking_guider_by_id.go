package bookingguiderhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) DeleteBookingGuiderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID := c.Param("id")

		bookingId, _ := strconv.Atoi(bookingID)

		err := hdl.bookingGuiderUC.DeleteBookingByID(c, bookingId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
