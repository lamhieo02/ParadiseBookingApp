package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) GetByID(ctx context.Context, id int) (*entities.Comment, error) {
	var data entities.Comment

	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
