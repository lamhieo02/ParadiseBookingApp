package placehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) GetPlaceByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		placeID := ctx.Param("id")
		id, err := strconv.Atoi(placeID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		place, err := hdl.placeUC.GetPlaceByID(ctx.Request.Context(), id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": place})
	}
}
