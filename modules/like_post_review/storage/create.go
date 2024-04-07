package likepostreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *likePostReviewStorage) Create(ctx context.Context, data *entities.LikePostReview) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
