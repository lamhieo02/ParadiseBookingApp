package reportstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *reportStorage) GetByID(ctx context.Context, id int) (*entities.Report, error) {
	var data entities.Report

	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
