package reportiomodel

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
}
