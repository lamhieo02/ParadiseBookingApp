package calendarguiderhandler

import (
	"net/http"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *calendarGuiderHandler) UpdateCalendarGuiderByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		calendarGuiderID, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		reqBody := calendarguideriomodel.UpdateCalendarGuiderReq{}
		if err := ctx.ShouldBind(&reqBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = hdl.calendarGuiderUC.UpdateCalendarGuiderByID(ctx, calendarGuiderID, &reqBody)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
