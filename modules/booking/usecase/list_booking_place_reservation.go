package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *bookingUseCase) ListPlaceReservationByVendor(ctx context.Context, vendorId, placeId int) ([]entities.Place, error) {

	if placeId != 0 {
		place, err := uc.PlaceSto.GetPlaceByID(ctx, placeId)
		if err != nil {
			return nil, err
		}

		if place == nil {
			return nil, common.ErrEntityNotFound(entities.Place{}.TableName(), errors.New("place not found"))
		}

		return []entities.Place{*place}, nil
	}

	conditions := []common.Condition{}
	conditions = append(conditions, common.Condition{
		Field:    "vendor_id",
		Operator: common.OperatorEqual,
		Value:    vendorId,
	})

	places, err := uc.PlaceSto.ListPlaceByCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}
	return places, nil
}
