package bookingguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) GetByID(ctx context.Context, id int) (*entities.BookingGuider, error) {
	var res entities.BookingGuider
	if err := s.db.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
