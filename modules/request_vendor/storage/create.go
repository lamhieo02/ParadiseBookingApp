package requestvendorstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *requestVendorSto) Create(ctx context.Context, data *entities.RequestVendor) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
