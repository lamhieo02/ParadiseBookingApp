package likepostreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *likePostReviewStorage) CountLikeByPostReview(ctx context.Context, postReviewID int) (*int64, error) {
	var count int64
	if err := s.db.Model(&entities.LikePostReview{}).Where("post_review_id = ? AND status = 1", postReviewID).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}
