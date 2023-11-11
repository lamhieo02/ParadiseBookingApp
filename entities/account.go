package entities

import (
	"paradise-booking/common"
)

type Account struct {
	common.SQLModel
	Username      string       `json:"username" gorm:"column:username"`
	Email         string       `json:"email" gorm:"column:email"`
	Name          string       `json:"name" gorm:"column:name"`
	AccountTypeId int          `json:"-" gorm:"column:account_type_id"`
	Status        int          `json:"-" gorm:"column:status"`
	Password      string       `json:"-" gorm:"column:password"`
	Address       string       `json:"address" gorm:"column:address"`
	Phone         string       `json:"phone" gorm:"column:phone"`
	Dob           string       `json:"Dob" gorm:"column:Dob"`
	Avatar        common.Image `json:"avatar" gorm:"avatar"`
}
