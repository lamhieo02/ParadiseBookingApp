package requestguiderhandler

import (
	"net/http"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *RequestGuiderHandler) CreateRequestGuider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dataReq *requestguideriomodel.CreateRequestGuiderReq
		if err := ctx.ShouldBind(&dataReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := hdl.requestGuiderUC.CreateRequestGuider(ctx, dataReq.ToEntity()); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
