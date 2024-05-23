package requestguiderhandler

import (
	"net/http"
	"paradise-booking/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *RequestGuiderHandler) ListRequestGuiderByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		// var filter *requestguideriomodel.Filter
		res, err := hdl.requestGuiderUC.ListByFilter(ctx, &paging, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
