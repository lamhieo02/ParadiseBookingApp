package postguideusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
)

type PostGuideCache interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type PostGuideSto interface {
	Create(ctx context.Context, data *entities.PostGuide) error
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
	DeleteByID(ctx context.Context, id int) error
	UpdateWithMap(ctx context.Context, id int, props map[string]interface{}) error
	UpdateByID(ctx context.Context, id int, postGuideData *entities.PostGuide) error
	ListByFilter(ctx context.Context, paging *common.Paging, filter *postguideiomodel.Filter) ([]*entities.PostGuide, error)
}

type AccountCache interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type postGuideUsecase struct {
	postGuideSto   PostGuideSto
	postGuideCache PostGuideCache
	accountCache   AccountCache
	googleMap      googlemapprovider.GoogleMap
}

func NewPostGuideUsecase(postGuideSto PostGuideSto, postGuideCache PostGuideCache, accountCache AccountCache, googleMap googlemapprovider.GoogleMap) *postGuideUsecase {
	return &postGuideUsecase{postGuideSto: postGuideSto, postGuideCache: postGuideCache, accountCache: accountCache, googleMap: googleMap}
}
