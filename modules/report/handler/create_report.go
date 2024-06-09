package reporthandler

import (
	"net/http"
	reportiomodel "paradise-booking/modules/report/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) CreateReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createBody reportiomodel.CreateReportReq
		if err := ctx.ShouldBind(&createBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := hdl.reportUseCase.CreateReport(ctx.Request.Context(), &createBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
