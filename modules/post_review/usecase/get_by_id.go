package postreviewusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/modules/post_review/convert"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (postReviewUsecase *postReviewUsecase) GetPostReviewByID(ctx context.Context, postReviewID int, accountID int) (*postreviewiomodel.PostReviewResp, error) {
	postReview, err := postReviewUsecase.postReviewStore.GetByID(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	comments, err := postReviewUsecase.commentStore.GetByPostReviewID(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	result := convert.ConvertPostReviewEntityToModelDetail(postReview, comments)

	likeCount, err := postReviewUsecase.likePostReviewSto.CountLikeByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	commentCount, err := postReviewUsecase.commentStore.CountCommentByPostReview(ctx, postReviewID)
	if err != nil {
		return nil, err
	}

	result.LikeCount = *likeCount
	result.CommentCount = *commentCount

	// get reply comment
	for i, v := range result.Comments {
		replyComments, err := postReviewUsecase.replyComment.GetBySourceCommentID(ctx, int(v.ID))
		if err != nil {
			return nil, err
		}

		tmpReplyComments := []postreviewiomodel.ReplyCommentResp{}
		for _, replyComment := range replyComments {
			tmpReplyComments = append(tmpReplyComments, postreviewiomodel.ReplyCommentResp{
				ID:        int64(replyComment.Id),
				Content:   replyComment.Content,
				Image:     replyComment.Image,
				Videos:    replyComment.Videos,
				AccountID: int64(replyComment.AccountID),
			})
		}

		result.Comments[i].ReplyComments = tmpReplyComments
	}

	// get location
	result.Country = postReview.Country
	result.State = postReview.State
	result.District = postReview.District

	likePostReview, err := postReviewUsecase.likePostReviewSto.FindDataByCondition(ctx, map[string]interface{}{
		"account_id":     accountID,
		"post_review_id": postReviewID,
	})

	if err != nil {
		return nil, err
	}

	if len(likePostReview) == 0 || likePostReview[0].Status == constant.UNLIKE_POST_REVIEW {
		result.IsLiked = false
	} else {
		result.IsLiked = true
	}

	return result, nil
}
