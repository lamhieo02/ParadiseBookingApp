package postreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewStorage) DeleteByID(ctx context.Context, postReviewID int) error {
	if err := s.db.Where("id = ?", postReviewID).Delete(&entities.PostReview{}).Error; err != nil {
		return err
	}

	return nil
}
