package bookingguiderusecase

import (
	"context"
	"paradise-booking/entities"
	bookingguiderconvert "paradise-booking/modules/booking_guider/convert"
	bookingguideriomodel "paradise-booking/modules/booking_guider/iomodel"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"sync"
)

func (uc *bookingGuiderUseCase) GetBookingByID(ctx context.Context, bookingGuiderID int) (*bookingguideriomodel.GetBookingGuiderResp, error) {

	wg := new(sync.WaitGroup)
	wg.Add(2)
	var (
		postGuide *postguideiomodel.GetPostGuideResp
		calendar  *entities.CalendarGuider
		err       error
	)

	data, err := uc.bookingGuiderSto.GetByID(ctx, bookingGuiderID)
	if err != nil {
		return nil, err
	}

	go func() {
		defer wg.Done()
		// get post guide info
		postGuide, err = uc.postGuideUC.GetPostGuideByID(ctx, data.PostGuideID)
		if err != nil {
			return
		}
	}()

	go func() {
		defer wg.Done()
		calendar, err = uc.calendarSto.GetByID(ctx, data.CalendarGuiderID)
		if err != nil {
			return
		}
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}

	res := bookingguiderconvert.ConvertBookingEntityToModel(data, postGuide)

	res.CalendarGuider.DateFrom = *calendar.DateFrom
	res.CalendarGuider.DateTo = *calendar.DateTo

	return res, nil
}
