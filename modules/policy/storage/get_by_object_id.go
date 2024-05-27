package policiesstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *policyStorage) GetByObjectID(ctx context.Context, objectId int, objectType int) ([]entities.Policy, error) {
	db := s.db.WithContext(ctx)

	var data []entities.Policy
	err := db.Table(entities.Policy{}.TableName()).Where("object_id = ? AND object_type = ?", objectId, objectType).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
