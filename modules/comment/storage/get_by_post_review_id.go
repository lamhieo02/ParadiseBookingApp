package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) GetByPostReviewID(ctx context.Context, postReviewID int) ([]*entities.Comment, error) {
	var comments []*entities.Comment

	if err := s.db.Where("post_review_id = ?", postReviewID).Order("id desc").Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}
