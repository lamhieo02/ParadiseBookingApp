package paymenthandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *paymentHandler) ListPaymentByVendorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		vendorID := ctx.Query("vendor_id")
		id, _ := strconv.Atoi(vendorID)

		payments, err := hdl.paymentUC.ListPaymentByVendorID(ctx, &paging, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": payments})
	}
}
