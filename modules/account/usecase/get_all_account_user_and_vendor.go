package accountusecase

import (
	"context"
	"paradise-booking/entities"
)

func (a *accountUseCase) GetAllAccountUserAndVendor(ctx context.Context) ([]entities.Account, error) {
	result, err := a.accountStorage.GetAllAccountUserAndVendor(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
