package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) DeleteByID(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&entities.PostGuide{}).Error; err != nil {
		return err
	}

	return nil
}
