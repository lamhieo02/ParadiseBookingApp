package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) ListAllPlace(ctx context.Context, paging *common.Paging, filter *iomodel.Filter) (result []iomodel.GetPlaceResp, err error) {

	paging.Process()
	places, err := uc.placeStorage.ListPlaces(ctx, paging, filter)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	// convert data to iomodel
	for _, place := range places {
		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place))
	}
	return result, nil
}
