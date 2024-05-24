package requestguiderhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *RequestGuiderHandler) ConfirmRequestGuider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestGuiderID := ctx.Query("request_guider_id")
		rgID, _ := strconv.Atoi(requestGuiderID)

		typeConfirm := ctx.Query("type")
		typeInt, _ := strconv.Atoi(typeConfirm)

		if err := hdl.requestGuiderUC.ConfirmRequestGuider(ctx, rgID, typeInt); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})

	}
}
