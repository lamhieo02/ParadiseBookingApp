package amenityhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *amenityHandler) ListAmenityByObjectId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		objectId := ctx.Query("object_id")
		objectType := ctx.Query("object_type")
		id, err := strconv.Atoi(objectId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		typeInt, err := strconv.Atoi(objectType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.amenityUC.ListAmenityByObjectID(ctx, id, typeInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})
	}
}
