package placehandler

import (
	"net/http"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *placeHandler) ListAllPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		places, err := hdl.placeUC.ListAllPlace(ctx.Request.Context(), &paging, nil)
		if err != nil {
			panic(err)
		}

		res := convert.ConvertPlaceToListModel(places, &paging)
		ctx.JSON(http.StatusOK, res)
	}
}
