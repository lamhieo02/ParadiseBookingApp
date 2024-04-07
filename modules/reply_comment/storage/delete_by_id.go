package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) DeleteByID(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&entities.ReplyComment{}).Error; err != nil {
		return err
	}

	return nil
}
