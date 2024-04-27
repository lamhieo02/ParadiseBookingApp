package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) Create(ctx context.Context, data *entities.PostGuide) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
