package policiesusecase

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/policy/iomodel"

	"gorm.io/gorm"
)

func (uc *policyUsecase) UpsertPolicy(ctx context.Context, dataReq *iomodel.CreatePolicyReq) error {

	for _, policy := range dataReq.Data.ListPolicy {
		data, err := uc.PolicyStore.GetByCondition(ctx, map[string]any{
			"object_id":       dataReq.Data.ObjectID,
			"object_type":     dataReq.Data.ObjectType,
			"group_policy_id": policy.GroupPolicyID,
		})

		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		// because place_id and group_policy_id is unique => len(data) must be 0 or 1
		if len(data) == 0 {
			// create new
			record := &entities.Policy{
				// PlaceId:       dataReq.Data.PlaceID,
				ObjectID:      dataReq.Data.ObjectID,
				ObjectType:    dataReq.Data.ObjectType,
				Name:          policy.Name,
				GroupPolicyId: policy.GroupPolicyID,
			}

			err = uc.PolicyStore.Create(ctx, record)
			if err != nil {
				return err
			}

		} else {
			if data[0].Name == policy.Name {
				continue
			}
			// update
			data[0].Name = policy.Name
			err := uc.PolicyStore.Update(ctx, &data[0])
			if err != nil {
				return err
			}
		}

	}
	return nil
}
