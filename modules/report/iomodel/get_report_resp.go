package reportiomodel

import "time"

type GetReportResp struct {
	ID          int      `json:"id"`
	ObjectID    int      `json:"object_id"`
	ObjectType  int      `json:"object_type"`
	ObjectName  string   `json:"object_name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	StatusID    int      `json:"status_id"`
	StatusName  string   `json:"status_name"`
	Videos      []string `json:"videos"`
	Images      []string `json:"images"`
	UserID      int      `json:"user_id"`
	User        User     `json:"user"`
}

type User struct {
	ID       int        `json:"id"`
	Role     string     `json:"role"`
	Email    string     `json:"email"`
	Username string     `json:"username"`
	FullName string     `json:"full_name"`
	Address  string     `json:"address"`
	Phone    string     `json:"phone"`
	DOB      string     `json:"dob"`
	Avt      string     `json:"avt"`
	Bio      string     `json:"bio"`
	Created  *time.Time `json:"created"`
	Updated  *time.Time `json:"updated"`
}
