package requestvendorstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *requestVendorSto) GetByUserID(ctx context.Context, userID int) (*entities.RequestVendor, error) {
	var requestVendor *entities.RequestVendor
	if err := s.db.Where("user_id = ?", userID).Find(&requestVendor).Error; err != nil {
		return nil, err
	}

	return requestVendor, nil
}
