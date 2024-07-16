package apiutils

import (
	"context"
	"paradise-booking/entities"
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

type apiUtilHandler struct {
	placeSto       PlaceStorage
	postGuideSto   PostGuideSto
	postGuideCache PostGuideCache
	placeCache     placeCache
}

func NewApiUtilsHandler(placeSto PlaceStorage, postguideSto PostGuideSto, postGuideCache PostGuideCache, placeCache placeCache) *apiUtilHandler {
	return &apiUtilHandler{placeSto: placeSto, postGuideSto: postguideSto, postGuideCache: postGuideCache, placeCache: placeCache}
}
