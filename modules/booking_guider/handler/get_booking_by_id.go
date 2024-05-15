package bookingguiderhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) GetBookingGuiderByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		bookingID, _ := c.GetQuery("id")

		bookingId, _ := strconv.Atoi(bookingID)

		res, err := hdl.bookingGuiderUC.GetBookingByID(c, bookingId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
