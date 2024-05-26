package entities

import "paradise-booking/common"

type Amenity struct {
	common.SQLModel
	ObjectID        int     `json:"object_id" gorm:"column:object_id"`
	ObjectType      int     `json:"object_type" gorm:"column:object_type"`
	Description     *string `json:"description" gorm:"column:description"`
	ConfigAmenityId int     `json:"config_amenity_id" gorm:"column:config_amenity_id"`
}

func (Amenity) TableName() string {
	return "amenities"
}

type ConfigAmenity struct {
	common.SQLModel
	Icon string `json:"icon" gorm:"column:icon"`
	Name string `json:"name" gorm:"column:name"`
	Type int    `json:"type" gorm:"column:type"`
}

func (ConfigAmenity) TableName() string {
	return "config_amenity"
}
