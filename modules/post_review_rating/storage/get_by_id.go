package postreviewratingstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewRatingStorage) GetByID(ctx context.Context, id int) (*entities.PostReviewRating, error) {
	var data entities.PostReviewRating
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
