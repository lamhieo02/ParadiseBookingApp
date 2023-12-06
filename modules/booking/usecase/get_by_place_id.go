package bookingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/booking/convert"
	"paradise-booking/modules/booking/iomodel"
)

func (uc *bookingUseCase) GetBookingByPlaceID(ctx context.Context, placeId int, paging *common.Paging) ([]iomodel.GetBookingByPlaceResp, error) {

	paging.Process()
	bookings, err := uc.bookingSto.GetByPlaceID(ctx, placeId, paging)
	if err != nil {
		return nil, err
	}

	var result []iomodel.GetBookingByPlaceResp

	// get place by id
	place, err := uc.PlaceSto.GetPlaceByID(ctx, placeId)
	if err != nil {
		return nil, common.ErrCannotGetEntity("place", err)
	}

	for _, booking := range bookings {

		// get account by id
		account, err := uc.AccountSto.GetProfileByID(ctx, booking.UserId)
		if err != nil {
			return nil, common.ErrCannotGetEntity("account", err)
		}

		result = append(result, *convert.ConvertBookingModelToGetByPlaceResp(account, &booking, place))

	}

	return result, nil
}
