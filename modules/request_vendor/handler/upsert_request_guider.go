package requestvendorhandler

import (
	"net/http"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *requestVendorHandler) UpsertRequestVendor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataReq *requestvendoriomodel.CreateRequestVendorReq
		if err := ctx.ShouldBind(&dataReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := hdl.requestVendorUC.UpsertRequestVendor(ctx, dataReq.ToEntity()); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
