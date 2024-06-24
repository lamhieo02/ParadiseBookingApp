package paymenthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *paymentHandler) UpdatePaymentStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paymentId := ctx.Query("id")
		id, _ := strconv.Atoi(paymentId)

		status := ctx.Query("status")
		statusID, _ := strconv.Atoi(status)

		if err := hdl.paymentUC.UpdateStatusPaymentByID(ctx, id, statusID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
