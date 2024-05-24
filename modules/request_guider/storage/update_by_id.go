package requestguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *RequestGuiderSto) UpdateByID(ctx context.Context, id int, data *entities.RequestGuider) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
