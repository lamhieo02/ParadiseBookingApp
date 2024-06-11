package reportiomodel

type Filter struct {
	ObjectID   int `json:"object_id"`
	ObjectType int `json:"object_type"`
	StatusID   int `json:"status_id"`
	UserID     int `json:"user_id"`
}
