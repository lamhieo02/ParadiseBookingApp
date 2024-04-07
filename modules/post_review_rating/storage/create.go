package postreviewratingstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewRatingStorage) Create(ctx context.Context, data *entities.PostReviewRating) error {

	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
