package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) UpdateByID(ctx context.Context, id int, data *entities.Comment) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
