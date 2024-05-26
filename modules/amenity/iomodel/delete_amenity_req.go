package iomodel

type DeleteAmenityReq struct {
	// IDPlace             int   `json:"place_id" form:"place_id"`
	ObjectID            int   `json:"object_id" form:"object_id"`
	ObjectType          int   `json:"object_type" form:"object_type"`
	ListConfigAmenityId []int `json:"list_config_amenity_id" form:"list_config_amenity_id"`
}
