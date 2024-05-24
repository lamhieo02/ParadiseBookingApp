package requestguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *RequestGuiderSto) GetByUserID(ctx context.Context, userID int) (*entities.RequestGuider, error) {
	var requestGuiders *entities.RequestGuider
	if err := s.db.Where("user_id = ?", userID).Find(&requestGuiders).Error; err != nil {
		return nil, err
	}

	return requestGuiders, nil
}
