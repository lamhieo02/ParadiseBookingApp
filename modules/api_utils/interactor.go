package apiutils

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	calendarguideriomodel "paradise-booking/modules/calendar_guider/iomodel"
)

type PlaceStorage interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error)
}

type PostGuideSto interface {
	ListPostGuideIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error)
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type PostGuideCache interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type placeCache interface {
	GetPlaceByID(ctx context.Context, placeId int) (*entities.Place, error)
}

type calendarGuider interface {
	ListByFilter(ctx context.Context, paging *common.Paging, filter *calendarguideriomodel.Filter) ([]*entities.CalendarGuider, error)
}

type postReviewSto interface {
	GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
}

type apiUtilHandler struct {
	placeSto          PlaceStorage
	postGuideSto      PostGuideSto
	postGuideCache    PostGuideCache
	placeCache        placeCache
	calendarGuiderSto calendarGuider
	postReviewSto     postReviewSto
}

func NewApiUtilsHandler(placeSto PlaceStorage, postguideSto PostGuideSto, postGuideCache PostGuideCache, placeCache placeCache, calendarGuiderSto calendarGuider, postReviewSto postReviewSto) *apiUtilHandler {
	return &apiUtilHandler{placeSto: placeSto, postGuideSto: postguideSto, postGuideCache: postGuideCache, placeCache: placeCache, calendarGuiderSto: calendarGuiderSto, postReviewSto: postReviewSto}
}
