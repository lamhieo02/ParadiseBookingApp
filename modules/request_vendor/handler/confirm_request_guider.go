package requestvendorhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *requestVendorHandler) ConfirmRequestVendor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestVendorID := ctx.Query("request_vendor_id")
		rvID, _ := strconv.Atoi(requestVendorID)

		typeConfirm := ctx.Query("type")
		typeInt, _ := strconv.Atoi(typeConfirm)

		if err := hdl.requestVendorUC.ConfirmRequestVendor(ctx, rvID, typeInt); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
