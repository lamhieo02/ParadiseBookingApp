package bookingusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/entities"
)

func (uc *bookingUseCase) ListPlaceReservationByVendor(ctx context.Context, vendorId, typeManage, placeId int) ([]entities.Place, error) {

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
	if typeManage == 1 { // type 1: is place available, have number_available >= 1
		conditions = append(conditions, common.Condition{
			Field:    "num_place_available",
			Operator: common.OperatorGreaterOrEqual,
			Value:    1,
		})
	}

	if typeManage == 2 { // type 2: is place not available, have number_available = 0
		conditions = append(conditions, common.Condition{
			Field:    "num_place_available",
			Operator: common.OperatorEqual,
			Value:    0,
		})
	}

	places, err := uc.PlaceSto.ListPlaceByCondition(ctx, conditions)
	if err != nil {
		return nil, err
	}
	return places, nil
}
