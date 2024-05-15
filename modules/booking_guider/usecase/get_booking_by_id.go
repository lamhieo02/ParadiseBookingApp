package bookingguiderusecase

import (
	"context"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

func (uc *bookingGuiderUseCase) GetBookingByID(ctx context.Context, bookingGuiderID int) (*bookingguideriomodel.GetBookingGuiderResp, error) {

	data, err := uc.bookingGuiderSto.GetByID(ctx, bookingGuiderID)
	if err != nil {
		return nil, err
	}

	res := bookingguiderconvert.ConvertBookingEntityToModel(data)
	calendar, err := uc.calendarSto.GetByID(ctx, data.CalendarGuiderID)
	if err != nil {
		return nil, err
	}

	res.CalendarGuider.DateFrom = *calendar.DateFrom
	res.CalendarGuider.DateTo = *calendar.DateTo

	return res, nil
}
