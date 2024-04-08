package postreviewusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/post_review/convert"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (postReviewUsecase *postReviewUsecase) ListPostReviewByAccountID(ctx context.Context, accountID int, paging *common.Paging) (*postreviewiomodel.ListPostReviewResp, error) {

	data, err := postReviewUsecase.postReviewStore.ListPostReviewByAccountID(ctx, accountID, paging)
	if err != nil {
		return nil, err
	}

	result := convert.ConvertListPostReviewToModel(data, paging)

	for i, v := range result.Data {
		likeCount, err := postReviewUsecase.likePostReviewSto.CountLikeByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		commentCount, err := postReviewUsecase.commentStore.CountCommentByPostReview(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		result.Data[i].LikeCount = *likeCount
		result.Data[i].CommentCount = *commentCount
	}

	return &result, nil
}
