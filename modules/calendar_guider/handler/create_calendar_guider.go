package calendarguiderhandler

import (
	"net/http"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"

	"github.com/gin-gonic/gin"
)

func (hdl *calendarGuiderHandler) CreateCalendarGuider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var createReq calendarguideriomodel.CreateCalendarGuiderReq
		if err := ctx.ShouldBind(&createReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := hdl.calendarGuiderUC.CreateCalendarGuider(ctx, &createReq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
