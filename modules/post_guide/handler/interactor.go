package postguidehandler

import (
	"context"
	"paradise-booking/common"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

type PostGuideUseCase interface {
	CreatePostGuide(ctx context.Context, data *postguideiomodel.CreatePostGuideReq) error
	GetPostGuideByID(ctx context.Context, id int) (*postguideiomodel.GetPostGuideResp, error)
	DeletePostGuideByID(ctx context.Context, id int) error
	UpdatePostGuideByID(ctx context.Context, id int, postGuideModel *postguideiomodel.UpdatePostGuideReq) error
	ListPostGuideByFilter(ctx context.Context, paging *common.Paging, filter *postguideiomodel.Filter) (*postguideiomodel.ListPostGuideResp, error)
}

type postGuideHandler struct {
	postGuideUC PostGuideUseCase
}

func NewPostGuideHandler(postGuideUC PostGuideUseCase) *postGuideHandler {
	return &postGuideHandler{postGuideUC: postGuideUC}
}
