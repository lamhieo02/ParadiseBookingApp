package requestguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *RequestGuiderSto) UpdateWithMap(ctx context.Context, data *entities.RequestGuider, props map[string]interface{}) error {
	db := s.db

	if err := db.Model(data).Updates(props).Error; err != nil {
		return err
	}
	return nil
}
