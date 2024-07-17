package bookinghandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/entities"
	"paradise-booking/modules/booking/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingHandler) ListBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("Account").(*entities.Account)

		var paging common.Paging
		var filter iomodel.FilterListBooking
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res, err := hdl.bookingUC.ListBooking(ctx.Request.Context(), &paging, &filter, requester.Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
