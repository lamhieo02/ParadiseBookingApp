package postreviewhandler

import (
	"context"
	"paradise-booking/common"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	postreviewratingiomodel "paradise-booking/modules/post_review_rating/iomodel"
)

type PostReviewUseCase interface {
	CreatePostReview(ctx context.Context, data *postreviewiomodel.CreatePostReviewReq) error
	UpdatePostReview(ctx context.Context, data *postreviewiomodel.UpdatePostReviewReq) error
	DeletePostReviewByID(ctx context.Context, postReviewID int) error
	GetPostReviewByID(ctx context.Context, postReviewID int, accountID int) (*postreviewiomodel.PostReviewResp, error)
	CommentPostReview(ctx context.Context, data *postreviewratingiomodel.CommentPostReviewRatingReq) error
	ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) (*postreviewiomodel.ListPostReviewResp, error)
	ListPostReviewByFilter(ctx context.Context, paging *common.Paging, filter *postreviewiomodel.Filter, accountID int64) (*postreviewiomodel.ListPostReviewResp, error)
}

type postReviewHandler struct {
	postReviewUC PostReviewUseCase
}

func NewPostReviewHandler(postReviewUC PostReviewUseCase) *postReviewHandler {
	return &postReviewHandler{postReviewUC: postReviewUC}
}
