package entities

import "paradise-booking/common"

type RequestGuider struct {
	common.SQLModel
	UserId       int    `gorm:"column:user_id" json:"user_id"`
	FullName     string `gorm:"column:full_name" json:"full_name"`
	Username     string `gorm:"column:username" json:"username"`
	Email        string `gorm:"column:email" json:"email"`
	Phone        string `gorm:"column:phone" json:"phone"`
	DOB          string `gorm:"column:dob" json:"dob"`
	Address      string `gorm:"column:address" json:"address"`
	Description  string `gorm:"column:description" json:"description"`
	Experience   string `gorm:"column:experience" json:"experience"`
	Reason       string `gorm:"column:reason" json:"reason"`
	GoalOfTravel string `gorm:"column:goal_of_travel" json:"goal_of_travel"`
	Languages    string `gorm:"column:languages" json:"languages"`
	Status       string `gorm:"column:status" json:"status"`
}

func (RequestGuider) TableName() string {
	return "request_guider"
}
