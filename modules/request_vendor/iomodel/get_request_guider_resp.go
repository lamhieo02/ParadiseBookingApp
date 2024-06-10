package requestvendoriomodel

import "time"

type GetRequestVendorResp struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	User        User   `json:"user"`
	FullName    string `json:"full_name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DOB         string `json:"dob"`
	Address     string `json:"address"`
	Description string `json:"description"`
	Experience  string `json:"experience"`
	Status      string `json:"status"`
}

type Filter struct {
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
