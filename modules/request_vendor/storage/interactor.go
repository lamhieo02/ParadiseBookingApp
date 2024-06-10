package requestvendorstorage

import "gorm.io/gorm"

type requestVendorSto struct {
	db *gorm.DB
}

func NewRequestVendorStorage(db *gorm.DB) *requestVendorSto {
	return &requestVendorSto{
		db: db,
	}
}
