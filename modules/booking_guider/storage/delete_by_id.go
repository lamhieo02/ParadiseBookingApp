package bookingguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) DeleteByID(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&entities.BookingGuider{}).Error; err != nil {
		return err
	}

	return nil
}
