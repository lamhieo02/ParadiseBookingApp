package calendarguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *calendarGuiderStorage) UpdateByID(ctx context.Context, id int, postGuideData *entities.CalendarGuider) error {
	if err := s.db.Where("id = ?", id).Updates(postGuideData).Error; err != nil {
		return err
	}

	return nil
}
