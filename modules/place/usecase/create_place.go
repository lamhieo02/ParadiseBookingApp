package placeusecase

import (
	"context"
	"paradise-booking/modules/place/convert"
	"paradise-booking/modules/place/iomodel"
)

func (uc *placeUseCase) CreatePlace(ctx context.Context, data *iomodel.CreatePlaceReq, emailVendor string) error {
	// convert iomodel to entity
	placeEntity := convert.ConvertPlaceCreateModelToEntity(data)

	// check vendor exist
	vendor, err := uc.accountSto.GetAccountByEmail(ctx, emailVendor)
	if err != nil {
		return err
	}

	placeEntity.VendorID = vendor.Id
	// create place
	if err := uc.placeStorage.Create(ctx, placeEntity); err != nil {
		return err
	}
	return nil
}
