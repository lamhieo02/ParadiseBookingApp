package bookingratinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingratinghandler) GetCommentByBookingID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		bookingID := ctx.Query("booking_id")
		id, err := strconv.Atoi(bookingID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		objectType := ctx.Query("object_type")
		objectTypeInt, err := strconv.Atoi(objectType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeRatingUC.GetCommentByBookingID(ctx, id, objectTypeInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
