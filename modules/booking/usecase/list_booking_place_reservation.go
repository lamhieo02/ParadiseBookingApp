package bookingusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *bookingUseCase) ListPlaceReservationByVendor(ctx context.Context, vendorId, typeManage int) (res []entities.Place, err error) {
	placeIdsBooked, err := uc.bookingSto.ListPlaceIds(ctx)
	if err != nil {
		return nil, err
	}

	if typeManage == 1 { // type 1: is place not reservation
		places, err := uc.PlaceSto.ListPlaceNotInIds(ctx, placeIdsBooked, vendorId)
		if err != nil {
			return nil, err
		}
		return places, nil
	}

	if typeManage == 2 {
		places, err := uc.PlaceSto.ListPlaceInIds(ctx, placeIdsBooked, vendorId)
		if err != nil {
			return nil, err
		}
		return places, nil
	}

	if typeManage == 3 {

		placesNotReservation, err := uc.PlaceSto.ListPlaceNotInIds(ctx, placeIdsBooked, vendorId)
		if err != nil {
			return nil, err
		}

		placesReservation, err := uc.PlaceSto.ListPlaceInIds(ctx, placeIdsBooked, vendorId)
		if err != nil {
			return nil, err
		}

		res = append(res, placesNotReservation...)
		res = append(res, placesReservation...)
		return res, nil
	}
	return nil, nil
}
