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
			// get owner of reply comment
			owner, err := postReviewUsecase.accountSto.GetProfileByID(ctx, int(replyComment.AccountID))
			if err != nil {
				return nil, err
			}
			tmpReplyComments = append(tmpReplyComments, postreviewiomodel.ReplyCommentResp{
				ID:        int64(replyComment.Id),
				Content:   replyComment.Content,
				Image:     replyComment.Image,
				Videos:    replyComment.Videos,
				AccountID: int64(replyComment.AccountID),
				Owner: postreviewiomodel.OwnerResp{
					UserName: owner.Username,
					Avatar:   owner.Avatar,
					FullName: owner.FullName,
				},
				DateComment: replyComment.CreatedAt,
			})
		}

		result.Comments[i].ReplyComments = tmpReplyComments
	}

	// get location
	result.Country = postReview.Country
	result.State = postReview.State
	result.District = postReview.District

	if accountID != 0 {
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
	}

	// get owner
	owner, err := postReviewUsecase.accountSto.GetProfileByID(ctx, int(result.PostOwnerID))
	if err != nil {
		return nil, err
	}
	result.PostOwner.Avatar = owner.Avatar
	result.PostOwner.FullName = owner.FullName
	result.PostOwner.UserName = owner.Username
	result.PostOwner.Email = owner.Email

	// get infor comments
	for i, v := range result.Comments {
		owner, err := postReviewUsecase.accountSto.GetProfileByID(ctx, int(v.AccountID))
		if err != nil {
			return nil, err
		}
		result.Comments[i].Owner.Avatar = owner.Avatar
		result.Comments[i].Owner.FullName = owner.FullName
		result.Comments[i].Owner.UserName = owner.Username
		result.Comments[i].Owner.Email = owner.Email
	}

	return result, nil
}
