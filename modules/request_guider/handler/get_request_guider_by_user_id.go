package requestguiderhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *RequestGuiderHandler) GetRequestGuiderByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		id, err := strconv.Atoi(userID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.requestGuiderUC.GetByUserID(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
