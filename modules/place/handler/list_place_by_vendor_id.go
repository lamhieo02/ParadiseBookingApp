package placehandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListPlaceByVendorID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("vendor_id")
		vendorID, err := strconv.Atoi(param)
		if err != nil {
			panic(err)
		}
		places, err := hdl.placeUC.ListPlaceByVendorByID(ctx.Request.Context(), vendorID)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, places)
	}
}
