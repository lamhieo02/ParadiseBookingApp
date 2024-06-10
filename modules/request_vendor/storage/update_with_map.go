package requestvendorstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *requestVendorSto) UpdateWithMap(ctx context.Context, data *entities.RequestVendor, props map[string]interface{}) error {
	db := s.db

	if err := db.Model(data).Updates(props).Error; err != nil {
		return err
	}
	return nil
}
