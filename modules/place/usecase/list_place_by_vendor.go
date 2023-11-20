package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) ListPlaceByVendor(ctx context.Context, vendorEmail string) (result []iomodel.GetPlaceResp, err error) {

	// get vendorID from vendorEmail
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, vendorEmail)
	if err != nil {
		return nil, common.ErrCannotGetEntity("account", err)
	}

	// get places by vendorID
	places, err := uc.placeStorage.ListPlaceByVendorID(ctx, vendor.Id)
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