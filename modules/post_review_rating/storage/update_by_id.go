package postreviewratingstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewRatingStorage) UpdateByID(ctx context.Context, id int, data *entities.PostReviewRating) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
