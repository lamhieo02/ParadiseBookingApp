package reportstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *reportStorage) Create(ctx context.Context, data *entities.Report) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
