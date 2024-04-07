package postreviewratingstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewRatingStorage) DeleteByID(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&entities.PostReviewRating{}).Error; err != nil {
		return err
	}

	return nil
}
