package requestguiderusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *requestGuiderUC) CreateRequestGuider(ctx context.Context, data *entities.RequestGuider) error {
	err := uc.requestGuiderSto.Create(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
