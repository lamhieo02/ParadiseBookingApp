package bookingratinghandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *bookingratinghandler) GetCommentByVendorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		vendorID := ctx.Query("vendor_id")
		id, err := strconv.Atoi(vendorID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		objectType := ctx.Query("object_type")
		objectTypeInt, err := strconv.Atoi(objectType)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res, err := hdl.placeRatingUC.GetCommentByVendorID(ctx, id, objectTypeInt)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res})

	}
}
