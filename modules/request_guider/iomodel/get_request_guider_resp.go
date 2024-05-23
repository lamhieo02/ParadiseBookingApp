package requestguideriomodel

type GetRequestGuiderResp struct {
	ID            int      `json:"id"`
	UserID        int      `json:"user_id"`
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
}

type Filter struct {
}
