package requestvendorstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *requestVendorSto) GetByID(ctx context.Context, id int) (*entities.RequestVendor, error) {
	var requestVendor *entities.RequestVendor
	if err := s.db.Where("id = ?", id).First(&requestVendor).Error; err != nil {
		return nil, err
	}

	return requestVendor, nil
}
