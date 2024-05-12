package bookingguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *bookingGuiderSto) UpdateStatus(ctx context.Context, bookingGuiderID int, status int) error {
	db := s.db
	booking := entities.BookingGuider{}
	if err := db.Table(booking.TableName()).Where("id = ?", bookingGuiderID).Update("status_id", status).Error; err != nil {
		return err
	}
	return nil
}
