package reportstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *reportStorage) UpdateByID(ctx context.Context, id int, data *entities.Report) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
