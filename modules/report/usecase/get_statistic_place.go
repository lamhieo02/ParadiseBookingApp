package reportusecase

import (
	"context"
	"errors"
	"fmt"
	"paradise-booking/common"
	"paradise-booking/constant"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"paradise-booking/utils"
	"time"

	"github.com/samber/lo"
)

func (uc *reportUseCase) GetStatisticsPlace(ctx context.Context, req reportiomodel.GetStatisticPlaceReq, vendorID int) (*reportiomodel.StatisticPlaceResp, error) {

	type valueBooking struct {
		BookingSuccess int
		BookingCancel  int
	}

	result := reportiomodel.StatisticPlaceResp{}
	mapColumnNameWithValuesBooking := map[string]valueBooking{}
	mapColumnNameWithValuesRevenue := map[string]float64{}
	layout := new(string)
	placeIds := []int{}

	columnsName, err := uc.getColumnNames(req.DateFrom, req.DateTo, req.Type, layout)
	if err != nil {
		return nil, err
	}

	lo.ForEach(columnsName, func(item string, _ int) {
		mapColumnNameWithValuesBooking[item] = valueBooking{BookingSuccess: 0, BookingCancel: 0}
		mapColumnNameWithValuesRevenue[item] = 0
	})

	if req.PlaceID != 0 {
		placeIds = append(placeIds, req.PlaceID)
	} else {
		// get all booking have status completed
		placeIds, err = uc.placeSto.ListPlaceIdsByCondition(ctx, 100, map[string]interface{}{
			"vendor_id": vendorID,
		})
		if err != nil {
			return nil, err
		}
	}

	dFrom := req.DateFrom + " 00:00:00"
	dTo := req.DateFrom + " 23:59:59"
	timeFrom, _ := utils.ParseStringToTimeWithHour(dFrom)
	timeTo, _ := utils.ParseStringToTimeWithHour(dTo)
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "place_id",
		Operator: common.OperatorIn,
		Value:    placeIds,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorIn,
		Value:    []int{constant.BookingStatusCompleted, constant.BookingStatusCancel},
	})
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: common.OperatorLessThanOrEqual,
		Value:    timeTo,
	})
	conditions = append(conditions, common.Condition{
		Field:    "created_at",
		Operator: common.OperatorGreaterOrEqual,
		Value:    timeFrom,
	})

	bookings, err := uc.bookingSto.ListAllBookingWithCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	for _, booking := range bookings {
		bookingDetail, err := uc.bookingDetailCache.GetByBookingID(ctx, booking.Id)
		if err != nil {
			return nil, err
		}

		keyMap := booking.CreatedAt.Format(*layout)

		if *layout == "1" {
			timeFrom, _ := utils.ParseStringToTime(req.DateFrom)
			timeTo := booking.CreatedAt

			// get the week number of timeTo when compare with timeFrom
			weekNumber := timeTo.AddDate(0, 0, 1).Sub(*timeFrom).Hours() / 24 / 7
			keyMap = fmt.Sprintf("Week %d", int(weekNumber)+1)
		}

		elementBooking := mapColumnNameWithValuesBooking[keyMap]
		elementRevenue := mapColumnNameWithValuesRevenue[keyMap]
		if booking.StatusId == constant.BookingStatusCompleted {
			result.TotalBookingSuccess++
			result.TotalRevenue += bookingDetail.TotalPrice
			elementBooking.BookingSuccess++
			elementRevenue += bookingDetail.TotalPrice
			mapColumnNameWithValuesRevenue[keyMap] = elementRevenue
			mapColumnNameWithValuesBooking[keyMap] = elementBooking
		}

		if booking.StatusId == constant.BookingStatusCancel {
			result.TotalBookingCancel++
			elementBooking.BookingCancel++
			mapColumnNameWithValuesBooking[keyMap] = elementBooking
		}
	}

	for _, columnName := range columnsName {
		mappingColumn := columnName
		if req.Type == constant.StatisticTypeWeek {
			timeFrom, _ := utils.ParseStringToTime(req.DateFrom)
			mappingNameColumns := uc.mappingColumnNameForTypeWeek(columnsName, *timeFrom)
			mappingColumn = mappingNameColumns[columnName]
		}

		result.StatisticBooking = append(result.StatisticBooking, reportiomodel.StatisticBooking{
			ColumnName:     mappingColumn,
			BookingSuccess: mapColumnNameWithValuesBooking[columnName].BookingSuccess,
			BookingCancel:  mapColumnNameWithValuesBooking[columnName].BookingCancel,
		})

		result.StatisticRevenue = append(result.StatisticRevenue, reportiomodel.StatisticRevenue{
			ColumnName: mappingColumn,
			Revenue:    mapColumnNameWithValuesRevenue[columnName],
		})
	}

	return &result, nil

}

func (uc *reportUseCase) mappingColumnNameForTypeWeek(columnsName []string, dateFrom time.Time) map[string]string {
	result := map[string]string{}
	// Week 1: 2021-01-01 - 2021-01-07
	// Week 2: 2021-01-08 - 2021-01-14
	// ...
	n := len(columnsName)
	for i := 0; i < n; i++ {
		weekDateFrom := dateFrom.Format("2006-01-02")
		weekDateTo := dateFrom.AddDate(0, 0, 6).Format("2006-01-02")
		result[columnsName[i]] = fmt.Sprintf("%s - %s", weekDateFrom, weekDateTo)

		dateFrom = dateFrom.AddDate(0, 0, 7)
	}

	return result
}

func (uc *reportUseCase) getColumnNames(timeFrom, timeTo string, _type int, layout *string) ([]string, error) {
	var result []string
	switch _type {
	case constant.StatisticTypeDay:
		// loop for all days in range: from req.DateFrom to req.DateTo
		*layout = "2006-01-02"
		dateFrom, _ := utils.ParseStringToTime(timeFrom)
		dateTo, _ := utils.ParseStringToTime(timeTo)
		for dateFrom.Before(dateTo.AddDate(0, 0, 1)) {
			result = append(result, dateFrom.Format("2006-01-02"))
			*dateFrom = dateFrom.AddDate(0, 0, 1)
		}
	case constant.StatisticTypeWeek:
		// loop for all weeks in range: from req.DateFrom to req.DateTo
		dateFrom, _ := utils.ParseStringToTime(timeFrom)
		dateTo, _ := utils.ParseStringToTime(timeTo)
		*layout = "1"
		cntWeek := 1
		for dateFrom.Before(dateTo.AddDate(0, 0, 1)) {
			result = append(result, fmt.Sprintf("Week %d", cntWeek))
			*dateFrom = dateFrom.AddDate(0, 0, 7)
			cntWeek++
		}
	case constant.StatisticTypeMonth:
		// loop for all months in range: from req.DateFrom to req.DateTo
		*layout = "2006-01"
		dateFrom, _ := utils.ParseStringToTime(timeFrom)
		dateTo, _ := utils.ParseStringToTime(timeTo)
		for dateFrom.Before(dateTo.AddDate(0, 0, 1)) {
			result = append(result, dateFrom.Format("2006-01"))
			*dateFrom = dateFrom.AddDate(0, 1, 0)
		}
	case constant.StatisticTypeYear:
		// loop for all years in range: from req.DateFrom to req.DateTo
		*layout = "2006"
		dateFrom, _ := utils.ParseStringToTime(timeFrom)
		dateTo, _ := utils.ParseStringToTime(timeTo)
		for dateFrom.Before(dateTo.AddDate(0, 0, 1)) {
			result = append(result, dateFrom.Format("2006"))
			*dateFrom = dateFrom.AddDate(1, 0, 0)
		}
	default:
		return nil, errors.New("invalid type")
	}

	return result, nil
}
