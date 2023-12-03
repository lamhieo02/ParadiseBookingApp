package bookinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) GetBookingByPlaceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeID := ctx.Query("place_id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingUC.GetBookingByPlaceID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, res)
	}
}
