package postguideusecase

import (
	"context"
	"paradise-booking/entities"
)

type PostGuideCache interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type PostGuideSto interface {
	Create(ctx context.Context, data *entities.PostGuide) error
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type postGuideUsecase struct {
	postGuideSto   PostGuideSto
	postGuideCache PostGuideCache
}

func NewPostGuideUsecase(postGuideSto PostGuideSto, postGuideCache PostGuideCache) *postGuideUsecase {
	return &postGuideUsecase{postGuideSto: postGuideSto, postGuideCache: postGuideCache}
}
