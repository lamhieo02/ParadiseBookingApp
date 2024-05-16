package bookingguiderhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) GetBookingGuiderByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")

		id, _ := strconv.Atoi(userID)

		res, err := hdl.bookingGuiderUC.GetBookingByUserID(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
