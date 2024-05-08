package entities

import (
	"paradise-booking/common"
	"time"
)

type CalendarGuider struct {
	common.SQLModel
	PostGuideId    int        `json:"post_guide_id" gorm:"column:post_guide_id"`
	GuiderId       int        `json:"guider_id" gorm:"column:guider_id"`
	Note           string     `json:"note" gorm:"column:note"`
	DateFrom       *time.Time `json:"date_from" gorm:"column:date_from"`
	DateTo         *time.Time `json:"date_to" gorm:"column:date_to"`
	PricePerPerson int        `json:"price_per_person" gorm:"column:price_per_person"`
	Status         bool       `json:"status" gorm:"column:status"`
	MaxGuest       int        `json:"max_guest" gorm:"column:max_guest"`
}

func (CalendarGuider) TableName() string {
	return "calendar_guider"
}
