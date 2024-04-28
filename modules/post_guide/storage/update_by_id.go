package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) UpdateByID(ctx context.Context, id int, postGuideData *entities.PostGuide) error {
	if err := s.db.Where("id = ?", id).Updates(postGuideData).Error; err != nil {
		return err
	}

	return nil
}
