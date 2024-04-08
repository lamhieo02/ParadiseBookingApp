package commentstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *commentStorage) CountCommentByPostReview(ctx context.Context, postReviewID int) (*int64, error) {
	var count int64
	if err := s.db.Model(&entities.Comment{}).Where("post_review_id = ?", postReviewID).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}
