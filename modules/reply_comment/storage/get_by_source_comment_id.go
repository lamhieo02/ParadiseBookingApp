package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) GetBySourceCommentID(ctx context.Context, sourceCommentID int) ([]entities.ReplyComment, error) {
	var data []entities.ReplyComment

	if err := s.db.Where("source_comment_id = ?", sourceCommentID).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
