package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) UpdateWithMap(ctx context.Context, id int, props map[string]interface{}) error {
	if err := s.db.Model(&entities.PostGuide{}).Where("id = ?", id).Updates(props).Error; err != nil {
		return err
	}

	return nil
}
