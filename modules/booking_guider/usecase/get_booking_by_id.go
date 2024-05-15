package bookingguiderusecase

import (
	"context"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
)

func (uc *bookingGuiderUseCase) GetBookingByID(ctx context.Context, bookingGuiderID int) error {

	data, err := uc.bookingGuiderSto.GetByID(ctx, bookingGuiderID)
	if err != nil {
		return err
	}

	res := bookingguiderconvert.ConvertBookingEntityToModel(data)
	calendar, err := uc.calendarSto.GetByID(ctx, data.CalendarGuiderID)
	if err != nil {
		return err
	}

	res.CalendarGuider.DateFrom = *calendar.DateFrom
	res.CalendarGuider.DateTo = *calendar.DateTo

	return nil
}
