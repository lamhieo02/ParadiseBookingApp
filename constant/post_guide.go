package constant

const (
	PostGuideTopicArtAndCulture = iota + 1
	PostGuideTopicEntertainment
	PostGuideTopicFoodAndDrink
	PostGuideTopicSport
	PostGuideTopicTours
	PostGuideTopicSightseeing
	PostGuideTopicWellness
)

var MapPostGuideTopic = map[int]string{
	PostGuideTopicArtAndCulture: "Art and Culture",
	PostGuideTopicEntertainment: "Entertainment",
	PostGuideTopicFoodAndDrink:  "Food and Drink",
	PostGuideTopicSport:         "Sport",
	PostGuideTopicTours:         "Tours",
	PostGuideTopicSightseeing:   "Sightseeing",
	PostGuideTopicWellness:      "Wellness",
}
