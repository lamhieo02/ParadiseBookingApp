package requestguideriomodel

type GetRequestGuiderResp struct {
	ID            int      `json:"id"`
	UserID        int      `json:"user_id"`
	User          User     `json:"user"`
	FullName      string   `json:"full_name"`
	Username      string   `json:"username"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	DOB           string   `json:"dob"`
	Address       string   `json:"address"`
	Description   string   `json:"description"`
	Experience    string   `json:"experience"`
	Reason        string   `json:"reason"`
	GoalsOfTravel []string `json:"goals_of_travel"`
	Languages     []string `json:"languages"`
	Status        string   `json:"status"`
}

type Filter struct {
}

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
}
