package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/booking/convert"
	"paradise-booking/modules/booking/iomodel"
)

func (uc *bookingUseCase) GetBookingByID(ctx context.Context, id int) (*iomodel.GetBookingResp, error) {
	booking, err := uc.bookingSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if booking == nil {
		return nil, common.ErrEntityNotFound("place", err)
	}

	// get account by id
	account, err := uc.AccountSto.GetProfileByID(ctx, booking.UserId)
	if err != nil {
		return nil, common.ErrCannotGetEntity("account", err)
	}

	// get place by id
	place, err := uc.PlaceSto.GetPlaceByID(ctx, booking.PlaceId)
	if err != nil {
		return nil, common.ErrCannotGetEntity("place", err)
	}

	result := convert.ConvertBookingModelToGetResp(account, booking, place)
	return result, nil
}
