package apiutils

import (
	"context"
	"paradise-booking/entities"
)

type PlaceStorage interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type PostGuideSto interface {
	ListPostGuideIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error)
}

type PostGuideCache interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type apiUtilHandler struct {
	placeSto       PlaceStorage
	postGuideSto   PostGuideSto
	postGuideCache PostGuideCache
}

func NewApiUtilsHandler(placeSto PlaceStorage, postguideSto PostGuideSto, postGuideCache PostGuideCache) *apiUtilHandler {
	return &apiUtilHandler{placeSto: placeSto, postGuideSto: postguideSto, postGuideCache: postGuideCache}
}
