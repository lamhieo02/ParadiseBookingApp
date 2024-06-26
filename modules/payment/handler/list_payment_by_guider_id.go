package paymenthandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *paymentHandler) ListPaymentByGuiderID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		guiderID := ctx.Query("guider_id")
		id, _ := strconv.Atoi(guiderID)

		bookingID := ctx.Query("booking_id")
		bookingIDInt, _ := strconv.Atoi(bookingID)

		payments, err := hdl.paymentUC.ListPaymentByGuiderID(ctx, &paging, id, bookingIDInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if bookingIDInt != 0 {
			paging.Page = 1
			paging.Limit = 1
			paging.Total = 1
		}

		ctx.JSON(http.StatusOK, gin.H{"data": payments, "paging": paging})
	}
}
