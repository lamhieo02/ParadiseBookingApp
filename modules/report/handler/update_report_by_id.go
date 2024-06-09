package reporthandler

import (
	"net/http"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) UpdateReportByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reportID := ctx.Param("id")
		id, _ := strconv.Atoi(reportID)

		var updateBody reportiomodel.UpdateReportReq
		if err := ctx.ShouldBind(&updateBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := hdl.reportUseCase.UpdateReportByID(ctx.Request.Context(), id, &updateBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
