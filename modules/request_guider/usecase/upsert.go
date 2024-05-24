package requestguiderusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (uc *requestGuiderUC) UpsertRequestGuider(ctx context.Context, data *entities.RequestGuider) error {
	// Check if request guider is already exist
	rgUser, err := uc.requestGuiderSto.GetByUserID(ctx, data.UserId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if rgUser.Id == 0 {
		data.Status = constant.RequestGuiderStatusProcessing
		err := uc.requestGuiderSto.Create(ctx, data)
		if err != nil {
			return err
		}
	} else {
		if err := uc.requestGuiderSto.UpdateByID(ctx, rgUser.Id, data); err != nil {
			return err
		}
	}

	return nil
}
