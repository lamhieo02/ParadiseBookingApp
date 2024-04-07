package postreviewratinghandler

import (
	"context"
	postreviewratingiomodel "paradise-booking/modules/post_review_rating/iomodel"
)

type PostReviewRatingUseCase interface {
	CommentPostReview(ctx context.Context, data *postreviewratingiomodel.CommentPostReviewRatingReq) error
}

type postReviewRatingHandler struct {
	postReviewRatingUC PostReviewRatingUseCase
}

func NewPostReviewRatingHandler(postReviewRatingUC PostReviewRatingUseCase) *postReviewRatingHandler {
	return &postReviewRatingHandler{postReviewRatingUC: postReviewRatingUC}
}
