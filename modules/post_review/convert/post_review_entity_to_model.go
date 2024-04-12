package convert

import (
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func ConvertPostReviewEntityToModel(postReviewEntity *entities.PostReview) *postreviewiomodel.PostReviewResp {
	return &postreviewiomodel.PostReviewResp{
		ID:          int64(postReviewEntity.Id),
		Title:       postReviewEntity.Title,
		TopicID:     postReviewEntity.Topic,
		TopicName:   constant.MapCategoryIDToName[postReviewEntity.Topic],
		PostOwnerID: int64(postReviewEntity.PostOwnerId),
		Content:     postReviewEntity.Content,
		Image:       postReviewEntity.Image,
		Lat:         postReviewEntity.Lat,
		Lng:         postReviewEntity.Lng,
		CreatedAt:   postReviewEntity.CreatedAt,
		UpdatedAt:   postReviewEntity.UpdatedAt,
		Country:     postReviewEntity.Country,
		State:       postReviewEntity.State,
		District:    postReviewEntity.District,
	}
}

func ConvertListPostReviewToModel(listPostReview []*entities.PostReview, paging *common.Paging) postreviewiomodel.ListPostReviewResp {
	var listPostReviewResp []postreviewiomodel.PostReviewResp
	for _, postReview := range listPostReview {
		listPostReviewResp = append(listPostReviewResp, *ConvertPostReviewEntityToModel(postReview))
	}
	return postreviewiomodel.ListPostReviewResp{
		Data:   listPostReviewResp,
		Paging: paging,
	}
}

func ConvertPostReviewEntityToModelDetail(postReviewEntity *entities.PostReview, comments []*entities.Comment) *postreviewiomodel.PostReviewResp {
	data := postreviewiomodel.PostReviewResp{
		ID:          int64(postReviewEntity.Id),
		Title:       postReviewEntity.Title,
		TopicID:     postReviewEntity.Topic,
		TopicName:   constant.MapCategoryIDToName[postReviewEntity.Topic],
		PostOwnerID: int64(postReviewEntity.PostOwnerId),
		Content:     postReviewEntity.Content,
		Image:       postReviewEntity.Image,
		Lat:         postReviewEntity.Lat,
		Lng:         postReviewEntity.Lng,
		CreatedAt:   postReviewEntity.CreatedAt,
		UpdatedAt:   postReviewEntity.UpdatedAt,
	}

	for _, comment := range comments {
		data.Comments = append(data.Comments, postreviewiomodel.CommentResp{
			ID:        int64(comment.Id),
			Content:   comment.Content,
			Image:     comment.Image,
			Videos:    comment.Videos,
			AccountID: comment.AccountID,
		})
	}

	return &data
}
