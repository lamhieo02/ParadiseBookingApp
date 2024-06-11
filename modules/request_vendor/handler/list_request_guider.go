package requestvendorhandler

import (
	"net/http"
	"paradise-booking/common"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *requestVendorHandler) ListRequestVendorByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter requestvendoriomodel.Filter
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		userID := ctx.Query("user_id")
		if userID != "" {
			userID, _ := strconv.Atoi(userID)
			filter.UserID = userID
		}

		status := ctx.Query("status")
		if status != "" {
			filter.Status = status
		}

		res, err := hdl.requestVendorUC.ListByFilter(ctx, &paging, &filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": res, "paging": paging})
	}
}
