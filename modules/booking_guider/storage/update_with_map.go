package bookingguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) UpdateWithMap(ctx context.Context, id int, props map[string]interface{}) error {
	if err := s.db.Model(&entities.BookingGuider{}).Where("id = ?", id).Updates(props).Error; err != nil {
		return err
	}

	return nil
}
