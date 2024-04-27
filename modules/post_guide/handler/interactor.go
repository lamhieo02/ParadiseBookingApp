package postguidehandler

import (
	"context"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

type PostGuideUseCase interface {
	CreatePostGuide(ctx context.Context, data *postguideiomodel.CreatePostGuideReq) error
	GetPostGuideByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type postGuideHandler struct {
	postGuideUC PostGuideUseCase
}

func NewPostGuideHandler(postGuideUC PostGuideUseCase) *postGuideHandler {
	return &postGuideHandler{postGuideUC: postGuideUC}
}
