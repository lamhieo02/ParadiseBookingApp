package bookingusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *bookingUseCase) ListPlaceNotReservationByVendor(ctx context.Context, vendorId int) (res []entities.Place, err error) {
	placeIdsBooked, err := uc.bookingSto.ListPlaceIds(ctx)
	if err != nil {
		return nil, err
	}

	places, err := uc.PlaceSto.ListPlaceNotInIds(ctx, placeIdsBooked, vendorId)
	if err != nil {
		return nil, err
	}
	return places, nil
}
