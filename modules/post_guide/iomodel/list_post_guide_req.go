package postguideiomodel

type Filter struct {
	PostOwnerId int     `json:"post_owner_id" form:"post_owner_id"`
	TopicID     int     `json:"topic_id" form:"topic_id"`
	Lat         float64 `json:"lat" form:"lat"`
	Lng         float64 `json:"lng" form:"lng"`
	State       string  `json:"state"`
}
