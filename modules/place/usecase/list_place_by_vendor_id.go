package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) ListPlaceByVendorByID(ctx context.Context, vendorID int) (result []iomodel.GetPlaceResp, err error) {

	// get places by vendorID
	places, err := uc.placeStorage.ListPlaceByVendorID(ctx, vendorID)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	if len(places) == 0 {
		return nil, common.ErrEntityNotFound("place", err)
	}

	// convert data to iomodel
	for _, place := range places {
		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place))
	}
	return result, nil
}
