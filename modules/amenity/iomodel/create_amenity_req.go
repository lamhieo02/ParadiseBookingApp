package iomodel

type CreateAmenityReq struct {
	// PlaceId           int             `json:"place_id"`
	ObjectID          int             `json:"object_id"`
	ObjectType        int             `json:"object_type"`
	ListDetailAmenity []DetailAmenity `json:"list_detail_amenity"`
}

type DetailAmenity struct {
	Description     *string `json:"description"`
	ConfigAmenityId int     `json:"config_amenity_id"`
}
