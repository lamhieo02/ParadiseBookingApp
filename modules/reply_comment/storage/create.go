package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) Create(ctx context.Context, data *entities.ReplyComment) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
