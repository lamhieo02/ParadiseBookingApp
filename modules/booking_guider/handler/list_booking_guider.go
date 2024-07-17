package bookingguiderhandler

import (
	"net/http"
	"paradise-booking/common"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingGuiderHandler) ListBookingGuider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter bookingguideriomodel.Filter
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.bookingGuiderUC.ListBooking(ctx.Request.Context(), &paging, &filter)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
