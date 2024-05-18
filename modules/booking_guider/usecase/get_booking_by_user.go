package bookingguiderusecase

import (
	"context"
	"paradise-booking/common"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

func (uc *bookingGuiderUseCase) GetBookingByUserID(ctx context.Context, userID int) ([]*bookingguideriomodel.GetBookingGuiderResp, error) {

	var res []*bookingguideriomodel.GetBookingGuiderResp

	condition := common.Condition{
		Field:    "user_id",
		Value:    userID,
		Operator: common.OperatorEqual,
	}
	data, err := uc.bookingGuiderSto.ListByCondition(ctx, []common.Condition{condition})
	if err != nil {
		return nil, err
	}

	for i, v := range data {
		res = append(res, bookingguiderconvert.ConvertBookingEntityToModel(v, nil))
		calendar, err := uc.calendarSto.GetByID(ctx, v.CalendarGuiderID)
		if err != nil {
			return nil, err
		}

		res[i].CalendarGuider.DateFrom = *calendar.DateFrom
		res[i].CalendarGuider.DateTo = *calendar.DateTo
	}

	return res, nil
}
