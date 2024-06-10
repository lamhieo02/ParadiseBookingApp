package requestvendorstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *requestVendorSto) UpdateByID(ctx context.Context, id int, data *entities.RequestVendor) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
