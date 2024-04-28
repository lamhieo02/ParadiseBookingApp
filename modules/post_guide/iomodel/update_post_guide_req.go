package postguideiomodel

type UpdatePostGuideReq struct {
	TopicID     int     `json:"topic_id" form:"topic_id"`
	Title       string  `json:"title" form:"title"`
	Description string  `json:"description" form:"description"`
	Cover       string  `json:"cover" form:"cover"`
	Lat         float64 `json:"lat" form:"lat"`
	Lng         float64 `json:"lng" form:"lng"`
}
