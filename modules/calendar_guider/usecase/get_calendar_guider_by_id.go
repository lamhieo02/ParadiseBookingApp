package calendarguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
	"paradise-booking/utils"

	"github.com/samber/lo"
)

func (uc *calendarGuiderUC) GetCalendarGuiderByID(ctx context.Context, id int) (*calendarguideriomodel.GetCalendarGuiderResp, error) {
	data, err := uc.calendarGuiderSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// handle to get max_guest available
	conditions := []common.Condition{
		{
			Field:    "calendar_guider_id",
			Operator: common.OperatorEqual,
			Value:    id,
		},
		{
			Field:    "status_id",
			Operator: common.OperatorIn,
			Value:    []int{constant.BookingGuiderStatusPending, constant.BookingGuiderStatusConfirmed},
		},
	}

	bookingGuiders, err := uc.bookingGuiderSto.ListByCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}

	numGuest := 0
	lo.ForEach(bookingGuiders, func(item *entities.BookingGuider, index int) {
		numGuest += item.NumberOfPeople
	})

	result := calendarguideriomodel.GetCalendarGuiderResp{
		ID:          data.Id,
		PostGuideID: data.GuiderId,
		GuiderID:    data.GuiderId,
		Note:        data.Note,
		DateFrom:    utils.ParseTimeWithHourToString(data.DateFrom),
		DateTo:      utils.ParseTimeWithHourToString(data.DateTo),
		Price:       data.PricePerPerson,
		Status:      data.Status,
		MaxGuest:    data.MaxGuest - numGuest,
	}

	return &result, nil
}
