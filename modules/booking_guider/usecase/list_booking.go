package bookingguiderusecase

import (
	"context"
	"paradise-booking/common"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
)

func (uc *bookingGuiderUseCase) ListBooking(ctx context.Context, paging *common.Paging, filter *bookingguideriomodel.Filter, userID int) ([]*bookingguideriomodel.GetBookingGuiderResp, error) {

	if paging != nil {
		paging.Process()
	}

	var res []*bookingguideriomodel.GetBookingGuiderResp

	data, err := uc.bookingGuiderSto.ListByFilter(ctx, paging, filter, userID)
	if err != nil {
		return nil, err
	}

	for i, v := range data {
		postGuide, err := uc.postGuideUC.GetPostGuideByID(ctx, v.PostGuideID)
		if err != nil {
			return nil, err
		}

		res = append(res, bookingguiderconvert.ConvertBookingEntityToModel(&v, postGuide))
		calendar, err := uc.calendarSto.GetByID(ctx, v.CalendarGuiderID)
		if err != nil {
			return nil, err
		}

		res[i].CalendarGuider.DateFrom = *calendar.DateFrom
		res[i].CalendarGuider.DateTo = *calendar.DateTo
	}

	return res, nil
}
