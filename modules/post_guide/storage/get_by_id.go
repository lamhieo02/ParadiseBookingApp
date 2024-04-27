package postguidestorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postGuideStorage) GetByID(ctx context.Context, id int) (*entities.PostGuide, error) {
	var postGuide entities.PostGuide
	if err := s.db.Where("id = ?", id).First(&postGuide).Error; err != nil {
		return nil, err
	}
	return &postGuide, nil
}
