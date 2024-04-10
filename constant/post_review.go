package constant

const (
	LIKE_POST_REVIEW   = 1
	UNLIKE_POST_REVIEW = 2
)

const (
	CATEGORY_DINNING        = 1
	CATEGORY_ENTERTAINMENT  = 2
	CATEGORY_ACCOMODATION   = 3
	CATEGORY_TRANSPORTATION = 4
	CATEGORY_SHOPPING       = 5
	CATEGORY_HEALTH         = 6
	CATEGORY_OTHERS         = 7
)

var MapCategoryIDToName = map[int]string{
	CATEGORY_DINNING:        "Dinning",
	CATEGORY_ENTERTAINMENT:  "Entertainment",
	CATEGORY_ACCOMODATION:   "Accomodation",
	CATEGORY_TRANSPORTATION: "Transportation",
	CATEGORY_SHOPPING:       "Shopping",
	CATEGORY_HEALTH:         "Health",
	CATEGORY_OTHERS:         "Others",
}
