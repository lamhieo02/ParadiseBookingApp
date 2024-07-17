package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
)

type PostReviewStore interface {
	Create(ctx context.Context, data *entities.PostReview) error
	UpdateWithMap(ctx context.Context, data *entities.PostReview, props map[string]interface{}) error
	GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
	UpdateByID(ctx context.Context, id int, data *entities.PostReview) error
	ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) ([]*entities.PostReview, error)
	DeleteByID(ctx context.Context, postReviewID int) error
	ListPostReviewByFilter(ctx context.Context, paging *common.Paging, filter *postreviewiomodel.Filter) ([]*entities.PostReview, error)
	ListByCondition(ctx context.Context, condition map[string]interface{}) ([]*entities.PostReview, error)
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
	FindDataByCondition(ctx context.Context, condition map[string]interface{}) ([]*entities.LikePostReview, error)
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
	googleMap         googlemapprovider.GoogleMap
}

func NewPostReviewUseCase(
	PostReviewStore PostReviewStore,
	commentSto CommentStorage,
	accountSto AccountStorage,
	likePostReviewSto LikePostReviewStorage,
	replyCommentSto ReplyCommentStorage,
	googleMap googlemapprovider.GoogleMap) *postReviewUsecase {
	return &postReviewUsecase{googleMap: googleMap, postReviewStore: PostReviewStore, accountSto: accountSto, commentStore: commentSto, likePostReviewSto: likePostReviewSto, replyComment: replyCommentSto}
}
