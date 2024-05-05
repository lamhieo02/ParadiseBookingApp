package calendarguiderstorage

import (
	"context"
	"paradise-booking/entities"
)

func (s *calendarGuiderStorage) GetByID(ctx context.Context, id int) (*entities.CalendarGuider, error) {
	db := s.db

	var data entities.CalendarGuider

	if err := db.Table(data.TableName()).Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
