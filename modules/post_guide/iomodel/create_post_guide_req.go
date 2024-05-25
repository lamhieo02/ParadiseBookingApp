package postguideiomodel

type CreatePostGuideReq struct {
	PostOwnerID int      `json:"post_owner_id" form:"post_owner_id" binding:"required"`
	TopicID     int      `json:"topic_id" form:"topic_id" binding:"required"`
	Title       string   `json:"title" form:"title" binding:"required"`
	Description string   `json:"description" form:"description" binding:"required"`
	Cover       string   `json:"cover" form:"cover" binding:"required"`
	Lat         float64  `json:"lat" form:"lat" binding:"required"`
	Lng         float64  `json:"lng" form:"lng" binding:"required"`
	Address     string   `json:"address" form:"address" binding:"required"`
	Languages   []string `json:"languages" form:"languages" binding:"required"`
}
