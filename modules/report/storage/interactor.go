package reportstorage

import "gorm.io/gorm"

type reportStorage struct {
	db *gorm.DB
}

func NewReportStorage(db *gorm.DB) *reportStorage {
	return &reportStorage{
		db: db,
	}
}
