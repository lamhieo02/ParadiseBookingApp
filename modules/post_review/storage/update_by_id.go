package postreviewstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *postReviewStorage) UpdateByID(ctx context.Context, id int, data *entities.PostReview) error {
	if err := s.db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
