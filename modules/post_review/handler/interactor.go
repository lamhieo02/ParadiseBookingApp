package postreviewhandler

import (
	"context"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

type PostReviewUseCase interface {
	CreatePostReview(ctx context.Context, data *postreviewiomodel.CreatePostReviewReq) error
}

type postReviewHandler struct {
	postReviewUC PostReviewUseCase
}

func NewPostReviewHandler(postReviewUC PostReviewUseCase) *postReviewHandler {
	return &postReviewHandler{postReviewUC: postReviewUC}
}
