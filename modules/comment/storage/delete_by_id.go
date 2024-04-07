package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) DeleteByID(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&entities.Comment{}).Error; err != nil {
		return err
	}

	return nil
}
