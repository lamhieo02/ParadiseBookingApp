package likepostreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *likePostReviewStorage) FindDataByCondition(ctx context.Context, condition map[string]interface{}) ([]*entities.LikePostReview, error) {
	var likePostReviews []*entities.LikePostReview

	if err := s.db.Where(condition).Find(&likePostReviews).Error; err != nil {
		return nil, err
	}
	return likePostReviews, nil
}
