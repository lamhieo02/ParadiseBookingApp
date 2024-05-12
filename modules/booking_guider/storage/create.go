package bookingguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) Create(ctx context.Context, data *entities.BookingGuider) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
