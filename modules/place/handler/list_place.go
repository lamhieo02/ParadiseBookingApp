package placehandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListAllPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paging := common.Paging{}
		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		places, err := hdl.placeUC.ListAllPlace(ctx.Request.Context(), &paging, nil)
		if err != nil {
			panic(err)
		}

		res := convert.ConvertPlaceToListModel(places, &paging)
		ctx.JSON(http.StatusOK, res)
	}
}
