package reporthandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) GetReportByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reportID := ctx.Param("id")
		id, _ := strconv.Atoi(reportID)

		data, err := hdl.reportUseCase.GetReportByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})
	}
}
