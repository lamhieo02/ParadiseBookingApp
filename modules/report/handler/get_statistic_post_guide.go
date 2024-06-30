package reporthandler

import (
	"net/http"
	"paradise-booking/common"
	reportiomodel "paradise-booking/modules/report/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) GetStatisticPostGuide() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("Account").(common.Requester)

		var req reportiomodel.GetStatisticPostGuideReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if valid, err := hdl.checkDateValid(req.DateFrom, req.DateTo, req.Type); !valid || err != nil {
			var msg interface{}
			msg = "invalid date range"
			if err != nil {
				msg = err
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		result, err := hdl.reportUseCase.GetStatisticsPostGuide(ctx, &req, requester.GetID())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}
