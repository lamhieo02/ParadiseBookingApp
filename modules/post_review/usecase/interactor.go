package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
)

type PostReviewStore interface {
	Create(ctx context.Context, data *entities.PostReview) error
	UpdateWithMap(ctx context.Context, data *entities.PostReview, props map[string]interface{}) error
	GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
	UpdateByID(ctx context.Context, id int, data *entities.PostReview) error
	ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error)
	DeleteByID(ctx context.Context, postReviewID int) error
}

type CommentStorage interface {
	Create(ctx context.Context, data *entities.Comment) error
	UpdateByID(ctx context.Context, id int, data *entities.Comment) error
	DeleteByID(ctx context.Context, id int) error
	GetByPostReviewID(ctx context.Context, postReviewID int) ([]*entities.Comment, error)
	CountCommentByPostReview(ctx context.Context, postReviewID int) (*int64, error)
}

type AccountStorage interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type LikePostReviewStorage interface {
	CountLikeByPostReview(ctx context.Context, postReviewID int) (*int64, error)
}

type ReplyCommentStorage interface {
	GetBySourceCommentID(ctx context.Context, sourceCommentID int) ([]entities.ReplyComment, error)
}

type postReviewUsecase struct {
	postReviewStore   PostReviewStore
	accountSto        AccountStorage
	commentStore      CommentStorage
	likePostReviewSto LikePostReviewStorage
	replyComment      ReplyCommentStorage
}

func NewPostReviewUseCase(
	PostReviewStore PostReviewStore,
	commentSto CommentStorage,
	accountSto AccountStorage,
	likePostReviewSto LikePostReviewStorage,
	replyCommentSto ReplyCommentStorage) *postReviewUsecase {
	return &postReviewUsecase{postReviewStore: PostReviewStore, accountSto: accountSto, commentStore: commentSto, likePostReviewSto: likePostReviewSto, replyComment: replyCommentSto}
}
