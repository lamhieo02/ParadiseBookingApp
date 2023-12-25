package placeusecase

import (
	"context"
	"log"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) ListPlaceByVendorByID(ctx context.Context, vendorID int, paging *common.Paging) (result []iomodel.GetPlaceResp, err error) {

	paging.Process()
	// get places by vendorID
	places, err := uc.placeStorage.ListPlaceByVendorID(ctx, vendorID, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity("place", err)
	}

	if len(places) == 0 {
		log.Printf("Not found any place by vendorID: %d", vendorID)
	}

	// convert data to iomodel
	if len(places) == 0 {
		return []iomodel.GetPlaceResp{}, nil
	}

	defaulRating := 0.0
	for _, place := range places {
		result = append(result, *convert.ConvertPlaceEntityToGetModel(&place, false, &defaulRating))
	}
	return result, nil
}
