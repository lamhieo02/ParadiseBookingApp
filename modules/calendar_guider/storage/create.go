package calendarguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *calendarGuiderStorage) Create(ctx context.Context, data *entities.CalendarGuider) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
