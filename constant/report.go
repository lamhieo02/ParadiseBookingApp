package constant

const (
	Report_Object_Type_User    = 1
	Report_Object_Type_Vendor  = 2
	Report_Object_Type_Guider  = 3
	Report_Object_Type_Place   = 4
	Report_Object_Type_Tour    = 5
	Report_Object_Type_Comment = 6
)

var MapReportObjectType = map[int]string{
	Report_Object_Type_User:    "user",
	Report_Object_Type_Vendor:  "vendor",
	Report_Object_Type_Guider:  "guider",
	Report_Object_Type_Place:   "place",
	Report_Object_Type_Tour:    "tour",
	Report_Object_Type_Comment: "comment",
}

const (
	Report_Status_Processing = 1
	Report_Status_Complete   = 2
)

var MapReportStatus = map[int]string{
	Report_Status_Processing: "processing",
	Report_Status_Complete:   "complete",
}
