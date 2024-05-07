package calendarguiderhandler

import (
	"net/http"
	"paradise-booking/common"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (hdl *calendarGuiderHandler) ListCalendarGuiderByFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter calendarguideriomodel.Filter

		page, _ := strconv.Atoi(ctx.Query("page"))
		limit, _ := strconv.Atoi(ctx.Query("limit"))

		paging.Page = page
		paging.Limit = limit

		// HANDLE FILTER
		postGuideID := ctx.Query("post_guide_id")
		if postGuideID != "" {
			postGuide, _ := strconv.Atoi(postGuideID)
			filter.PostGuideID = postGuide
		}

		guiderID := ctx.Query("guider_id")
		if guiderID != "" {
			guider, _ := strconv.Atoi(guiderID)
			filter.GuiderID = guider
		}

		dateFrom := ctx.Query("date_from")
		dateTo := ctx.Query("date_to")
		filter.DateFrom = dateFrom
		filter.DateTo = dateTo

		pricePerPerson := ctx.Query("price_per_person")
		if pricePerPerson != "" {
			price, _ := strconv.ParseFloat(pricePerPerson, 64)
			filter.PricePerPerson = price
		}

		status := ctx.Query("status")
		if status != "" {
			statusBool, _ := strconv.ParseBool(status)
			filter.Status = &statusBool
		}

		data, err := hdl.calendarGuiderUC.ListCalendarGuiderByFilter(ctx, &paging, &filter)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": data})

	}
}
