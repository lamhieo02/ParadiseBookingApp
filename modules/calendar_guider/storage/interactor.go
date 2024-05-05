package calendarguiderstorage

import "gorm.io/gorm"

type calendarGuiderStorage struct {
	db *gorm.DB
}

func NewCalendarGuiderStorage(db *gorm.DB) *calendarGuiderStorage {
	return &calendarGuiderStorage{
		db: db,
	}
}
