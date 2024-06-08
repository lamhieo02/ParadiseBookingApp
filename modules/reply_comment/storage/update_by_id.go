package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) UpdateByID(ctx context.Context, id int, data *entities.ReplyComment) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
