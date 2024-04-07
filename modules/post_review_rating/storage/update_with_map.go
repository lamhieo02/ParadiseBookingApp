package postreviewratingstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewRatingStorage) UpdateWithMap(ctx context.Context, data *entities.PostReviewRating, props map[string]interface{}) error {
	if err := s.db.Model(&data).Updates(props).Error; err != nil {
		return err
	}
	return nil
}
