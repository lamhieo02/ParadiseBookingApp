package reporthandler

import (
	"net/http"
	"paradise-booking/common"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) ListReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter reportiomodel.Filter
		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		objectID := ctx.Query("object_id")
		if objectID != "" {
			objectID, _ := strconv.Atoi(objectID)
			filter.ObjectID = objectID
		}

		objectType := ctx.Query("object_type")
		if objectType != "" {
			objectType, _ := strconv.Atoi(objectType)
			filter.ObjectType = objectType
		}

		statusID := ctx.Query("status_id")
		if statusID != "" {
			statusID, _ := strconv.Atoi(statusID)
			filter.StatusID = statusID
		}

		userID := ctx.Query("user_id")
		if userID != "" {
			userID, _ := strconv.Atoi(userID)
			filter.UserID = userID
		}

		result, err := hdl.reportUseCase.ListReport(ctx.Request.Context(), &paging, &filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}
