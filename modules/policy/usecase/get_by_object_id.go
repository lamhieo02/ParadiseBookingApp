package policiesusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *policyUsecase) GetPolicyByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.Policy, error) {

	data, err := uc.PolicyStore.GetByObjectID(ctx, objectID, objectType)
	if err != nil {
		return nil, err
	}

	return data, nil
}
