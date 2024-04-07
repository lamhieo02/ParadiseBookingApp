package replycommentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *replyCommentStorage) GetByID(ctx context.Context, id int) (*entities.ReplyComment, error) {
	var data entities.ReplyComment

	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
