package entities

import "paradise-booking/common"

type Report struct {
	common.SQLModel
}

func (Report) TableName() string {
	return "reports"
}
