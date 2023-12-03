package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/booking/convert"
	"paradise-booking/modules/booking/iomodel"
)

func (uc *bookingUseCase) GetBookingByPlaceID(ctx context.Context, placeId int) ([]iomodel.GetBookingByPlaceResp, error) {
	bookings, err := uc.bookingSto.GetByPlaceID(ctx, placeId)
	if err != nil {
		return nil, err
	}

	var result []iomodel.GetBookingByPlaceResp

	for _, booking := range bookings {

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

		result = append(result, *convert.ConvertBookingModelToGetByPlaceResp(account, &booking, place))

	}

	return result, nil
}
