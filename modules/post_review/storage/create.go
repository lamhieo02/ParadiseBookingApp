package postreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewStorage) Create(ctx context.Context, data *entities.PostReview) error {

	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
