package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewratingiomodel "paradise-booking/modules/post_review_rating/iomodel"
)

func (postReviewUC *postReviewUsecase) CommentPostReview(ctx context.Context, data *postreviewratingiomodel.CommentPostReviewRatingReq) error {

	// check account is exist
	account, err := postReviewUC.accountSto.GetProfileByID(ctx, int(data.AccountID))
	if err != nil {
		return err
	}

	if account == nil {
		return common.ErrEntityNotFound(entities.Account{}.TableName(), nil)
	}

	commentEntity := &entities.Comment{
		Content:      data.Comment,
		Image:        data.Image,
		Videos:       data.Videos,
		AccountID:    data.AccountID,
		PostReviewID: int64(data.PostReviewID),
	}

	// create comment
	if err := postReviewUC.commentStore.Create(ctx, commentEntity); err != nil {
		return err
	}

	return nil
}
