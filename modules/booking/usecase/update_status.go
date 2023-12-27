package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
)

func (uc *bookingUseCase) UpdateStatusBooking(ctx context.Context, bookingID, status int) error {
	// update status booking
	if err := uc.bookingSto.UpdateStatus(ctx, bookingID, status); err != nil {
		return common.ErrCannotUpdateEntity(entities.Booking{}.TableName(), err)
	}

	if status == constant.BookingStatusCompleted || status == constant.BookingStatusCancel {
		booking, err := uc.bookingSto.GetByID(ctx, bookingID)
		if err != nil {
			return common.ErrCannotGetEntity(entities.Booking{}.TableName(), err)
		}

		place, err := uc.PlaceSto.GetPlaceByID(ctx, booking.PlaceId)
		if err != nil {
			return common.ErrCannotGetEntity(entities.Place{}.TableName(), err)
		}

		newNumPlace := place.NumPlaceAvailable + 1
		if err := uc.PlaceSto.UpdateWithMap(ctx, place, map[string]interface{}{"num_place_available": newNumPlace}); err != nil {
			return common.ErrCannotUpdateEntity(entities.Place{}.TableName(), err)
		}
	}

	return nil
}
