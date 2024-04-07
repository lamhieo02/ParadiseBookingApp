package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) Create(ctx context.Context, data *entities.Comment) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
