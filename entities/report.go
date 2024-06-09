package entities

import "paradise-booking/common"

type Report struct {
	common.SQLModel
	ObjectID    int    `json:"object_id" gorm:"column:object_id"`
	ObjectType  int    `json:"object_type" gorm:"column:object_type"`
	Type        string `json:"type" gorm:"column:type"`
	Description string `json:"description" gorm:"column:description"`
	StatusID    int    `json:"status_id" gorm:"column:status_id"`
	Videos      string `json:"videos" gorm:"column:videos"`
	Images      string `json:"images" gorm:"column:images"`
}

func (Report) TableName() string {
	return "reports"
}
