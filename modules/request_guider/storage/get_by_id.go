package requestguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *RequestGuiderSto) GetByID(ctx context.Context, id int) (*entities.RequestGuider, error) {
	var requestGuider *entities.RequestGuider
	if err := s.db.Where("id = ?", id).First(&requestGuider).Error; err != nil {
		return nil, err
	}

	return requestGuider, nil
}
