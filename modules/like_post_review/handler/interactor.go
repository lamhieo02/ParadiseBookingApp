package likepostreviewhandler

import (
	"context"
	likepostreviewiomodel "paradise-booking/modules/like_post_review/iomodel"
)

type LikePostReviewUseCase interface {
	LikePostReview(ctx context.Context, req *likepostreviewiomodel.LikePostReviewReq) error
}

type likePostReviewHandler struct {
	likePostReview LikePostReviewUseCase
}

func NewLikePostReviewHandler(likePostReview LikePostReviewUseCase) *likePostReviewHandler {
	return &likePostReviewHandler{likePostReview: likePostReview}
}
