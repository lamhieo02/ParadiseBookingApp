package requestvendorusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (uc *requestVendorUC) UpsertRequestVendor(ctx context.Context, data *entities.RequestVendor) error {
	// Check if request vendor is already exist
	rgUser, err := uc.requestVendorSto.GetByUserID(ctx, data.UserId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if rgUser.Id == 0 {
		data.Status = constant.RequestVendorStatusProcessing
		err := uc.requestVendorSto.Create(ctx, data)
		if err != nil {
			return err
		}
	} else {
		if err := uc.requestVendorSto.UpdateByID(ctx, rgUser.Id, data); err != nil {
			return err
		}
	}

	return nil
}
