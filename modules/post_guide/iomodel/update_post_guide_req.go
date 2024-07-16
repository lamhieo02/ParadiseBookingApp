package postguideiomodel

type UpdatePostGuideReq struct {
	TopicID     int      `json:"topic_id" form:"topic_id"`
	Title       string   `json:"title" form:"title"`
	Description string   `json:"description" form:"description"`
	Images      []string `json:"images" form:"images"`
	Lat         float64  `json:"lat" form:"lat"`
	Lng         float64  `json:"lng" form:"lng"`
	Address     string   `json:"address" form:"address"`
	Languages   []string `json:"languages" form:"languages"`
	Schedule    string   `json:"schedule" form:"schedule"`
	Country     string   `json:"country"`
	State       string   `json:"state"`
	District    string   `json:"district"`
}
