package calendarguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *calendarGuiderStorage) DeleteByID(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(entities.CalendarGuider{}.TableName()).Where("id = ?", id).Delete(&entities.CalendarGuider{}).Error; err != nil {
		return err
	}

	return nil
}
