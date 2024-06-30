package reportusecase

import (
	"context"
	"fmt"
	"paradise-booking/common"
	"paradise-booking/constant"
	reportiomodel "paradise-booking/modules/report/iomodel"
	"paradise-booking/utils"

	"github.com/samber/lo"
)

func (uc *reportUseCase) GetStatisticsPostGuide(ctx context.Context, req *reportiomodel.GetStatisticPostGuideReq, postOwnerID int) (*reportiomodel.StatisticPostGuideResp, error) {

	type valueBooking struct {
		BookingSuccess int
		BookingCancel  int
	}

	result := reportiomodel.StatisticPostGuideResp{}
	mapColumnNameWithValuesBooking := map[string]valueBooking{}
	mapColumnNameWithValuesRevenue := map[string]float64{}
	layout := new(string)
	postGuideIds := []int{}

	columnsName, err := uc.getColumnNames(req.DateFrom, req.DateTo, req.Type, layout)
	if err != nil {
		return nil, err
	}

	lo.ForEach(columnsName, func(item string, _ int) {
		mapColumnNameWithValuesBooking[item] = valueBooking{BookingSuccess: 0, BookingCancel: 0}
		mapColumnNameWithValuesRevenue[item] = 0
	})

	if req.PostGuideID != 0 {
		postGuideIds = append(postGuideIds, req.PostGuideID)
	} else {
		// get all booking have status completed
		postGuideIds, err = uc.postGuideSto.ListPostGuideIdsByCondition(ctx, 100, map[string]interface{}{
			"post_owner_id": postOwnerID,
		})
		if err != nil {
			return nil, err
		}
	}

	timeFrom, _ := utils.ParseStringToTime(req.DateFrom)
	timeTo, _ := utils.ParseStringToTime(req.DateTo)
	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "post_guide_id",
		Operator: common.OperatorIn,
		Value:    postGuideIds,
	})
	conditions = append(conditions, common.Condition{
		Field:    "status_id",
		Operator: common.OperatorIn,
		Value:    []int{constant.BookingGuiderStatusCompleted, constant.BookingGuiderStatusCancel},
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

	bookings, err := uc.bookingGuiderSto.ListByCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	for _, booking := range bookings {
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
		if booking.StatusID == constant.BookingGuiderStatusCompleted {
			result.TotalBookingSuccess++
			result.TotalRevenue += booking.TotalPrice
			elementBooking.BookingSuccess++
			elementRevenue += booking.TotalPrice
			mapColumnNameWithValuesRevenue[keyMap] = elementRevenue
			mapColumnNameWithValuesBooking[keyMap] = elementBooking
		}

		if booking.StatusID == constant.BookingGuiderStatusCancel {
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
