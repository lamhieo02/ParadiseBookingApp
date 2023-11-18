package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) GetPlaceByID(ctx context.Context, placeID int) (result *iomodel.GetPlaceResp, err error) {
	place, err := uc.placeStorage.GetPlaceByID(ctx, placeID)
	if err != nil {
		return nil, err
	}

	if place == nil {
		return nil, common.ErrEntityNotFound("place", err)
	}

	result = convert.ConvertPlaceEntityToGetModel(place)
	return result, nil
}
