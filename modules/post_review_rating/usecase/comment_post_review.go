package postreviewratingusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewratingiomodel "paradise-booking/modules/post_review_rating/iomodel"
)

func (postReviewRatingUsecase *postReviewRatingUsecase) CommentPostReview(ctx context.Context, data *postreviewratingiomodel.CommentPostReviewRatingReq) error {

	// check account is exist
	account, err := postReviewRatingUsecase.accountSto.GetProfileByID(ctx, int(data.AccountID))
	if err != nil {
		return err
	}

	if account == nil {
		return common.ErrEntityNotFound(entities.Account{}.TableName(), nil)
	}

	commentEntity := &entities.Comment{
		Content:   data.Comment,
		Image:     data.Image,
		Videos:    data.Videos,
		AccountID: data.AccountID,
	}

	// create comment
	if err := postReviewRatingUsecase.commentStore.Create(ctx, commentEntity); err != nil {
		return err
	}

	postReviewEntity := &entities.PostReviewRating{
		PostReviewId: int(data.PostReviewID),
		CommentID:    commentEntity.Id,
	}

	// create post review rating
	if err := postReviewRatingUsecase.postReviewRatingStore.Create(ctx, postReviewEntity); err != nil {
		return err
	}

	return nil
}
