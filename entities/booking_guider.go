package entities

import (
	"paradise-booking/common"
)

type BookingGuider struct {
	common.SQLModel
	CalendarGuiderID int     `json:"calendar_guider_id" gorm:"column:calendar_guider_id"`
	Email            string  `json:"email" gorm:"column:email"`
	NumberOfPeople   int     `json:"number_of_people" gorm:"column:number_of_people"`
	Name             string  `json:"name" gorm:"column:name"`
	Note             string  `json:"note" gorm:"column:note"`
	StatusID         int     `json:"status_id" gorm:"column:status_id"`
	TotalPrice       float64 `json:"total_price" gorm:"column:total_price"`
	Phone            string  `json:"phone" gorm:"column:phone"`
	PaymentMethod    int     `json:"payment_method" gorm:"column:payment_method"`
	UserID           int     `json:"user_id" gorm:"column:user_id"`
	PostGuideID      int     `json:"post_guide_id" gorm:"column:post_guide_id"`
	GuiderID         int     `json:"guider_id" gorm:"column:guider_id"`
}

func (BookingGuider) TableName() string {
	return "booking_guider"
}
