package requestguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *RequestGuiderSto) Create(ctx context.Context, data *entities.RequestGuider) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
