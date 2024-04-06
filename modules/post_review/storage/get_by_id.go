package postreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewStorage) GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error) {
	var postReview entities.PostReview
	if err := s.db.Where("id = ?", postReviewID).First(&postReview).Error; err != nil {
		return nil, err
	}

	return &postReview, nil
}
