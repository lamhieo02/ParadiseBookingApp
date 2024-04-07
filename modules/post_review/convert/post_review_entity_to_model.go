package convert

import (
	"paradise-booking/common"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func ConvertPostReviewEntityToModel(postReviewEntity *entities.PostReview) *postreviewiomodel.PostReviewResp {
	return &postreviewiomodel.PostReviewResp{
		ID:          int64(postReviewEntity.Id),
		Title:       postReviewEntity.Title,
		Topic:       postReviewEntity.Topic,
		PostOwnerID: int64(postReviewEntity.PostOwnerId),
		Content:     postReviewEntity.Content,
		Image:       postReviewEntity.Image,
		Lat:         postReviewEntity.Lat,
		Lng:         postReviewEntity.Lng,
		CreatedAt:   postReviewEntity.CreatedAt,
		UpdatedAt:   postReviewEntity.UpdatedAt,
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
