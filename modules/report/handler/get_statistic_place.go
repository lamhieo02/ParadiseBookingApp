package reporthandler

import (
	"errors"
	"net/http"
	"paradise-booking/common"
	"paradise-booking/constant"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"paradise-booking/utils"

	"github.com/gin-gonic/gin"
)

func (hdl *reportHandler) GetStatisticPlace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requester := ctx.MustGet("Account").(common.Requester)

		var req reportiomodel.GetStatisticPlaceReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		if valid, err := hdl.checkDateValid(req.DateFrom, req.DateTo, req.Type); !valid || err != nil {
			var msg interface{}
			msg = "invalid date range"
			if err != nil {
				msg = err
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": msg})
			return
		}

		result, err := hdl.reportUseCase.GetStatisticsPlace(ctx, &req, requester.GetID())
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func (hdl *reportHandler) checkDateValid(timeFrom, timeTo string, _type int) (bool, error) {
	dateFrom, err := utils.ParseStringToTime(timeFrom)
	if err != nil {
		return false, err
	}

	dateTo, err := utils.ParseStringToTime(timeTo)
	if err != nil {
		return false, err
	}

	switch _type {
	case constant.StatisticTypeDay:
		// date range is 7 days
		if dateTo.Sub(*dateFrom).Hours() > 24*7 {
			return false, nil
		}
	case constant.StatisticTypeWeek:
		// date range is 12 weeks
		if dateTo.Sub(*dateFrom).Hours() > 24*7*12 {
			return false, nil
		}
	case constant.StatisticTypeMonth:
		// Calculate the difference in years and months
		yearDiff := dateTo.Year() - dateFrom.Year()
		monthDiff := int(dateTo.Month()) - int(dateFrom.Month())
		dayDiff := dateTo.Day() - dateFrom.Day()

		// Adjust for cases where monthDiff is negative or the day is earlier
		if monthDiff < 0 || (monthDiff == 0 && dayDiff < 0) {
			yearDiff--
			monthDiff += 12
		}

		// Check if the total difference is 12 months
		return yearDiff == 1 && monthDiff == 0 && dayDiff == 0, nil

	case constant.StatisticTypeYear:
		// date range is 5 years
		yearDiff := dateTo.Year() - dateFrom.Year()
		monthDiff := int(dateTo.Month()) - int(dateFrom.Month())
		dayDiff := dateTo.Day() - dateFrom.Day()

		// Adjust for cases where monthDiff is negative
		if monthDiff < 0 || (monthDiff == 0 && dayDiff < 0) {
			yearDiff--
		}

		// Check if the total difference is 5 years
		return yearDiff == 5 && (monthDiff > 0 || (monthDiff == 0 && dayDiff >= 0)), nil

	default:
		return false, errors.New("invalid type")
	}

	return true, nil
}
