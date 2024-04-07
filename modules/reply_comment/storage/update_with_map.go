package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) UpdateWithMap(ctx context.Context, data *entities.ReplyComment, props map[string]interface{}) error {
	if err := s.db.Model(&data).Updates(props).Error; err != nil {
		return err
	}
	return nil
}
