package entities

import "paradise-booking/common"

type Policy struct {
	common.SQLModel
	ObjectID      int    `json:"object_id" gorm:"column:object_id"`
	ObjectType    int    `json:"object_type" gorm:"column:object_type"`
	Name          string `json:"name" gorm:"column:name"`
	GroupPolicyId int    `json:"group_policy_id" gorm:"column:group_policy_id"`
}

func (Policy) TableName() string {
	return "policies"
}
