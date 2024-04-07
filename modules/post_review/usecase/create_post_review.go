package postreviewusecase

import (
	"context"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
)

func (postReviewUsecase *postReviewUsecase) CreatePostReview(ctx context.Context, data *postreviewiomodel.CreatePostReviewReq) error {
	models := entities.PostReview{
		PostOwnerId: int(data.AccountID),
		Title:       data.Title,
		Topic:       data.Topic,
		Content:     data.Content,
		Lat:         data.Lat,
		Lng:         data.Lng,
		Image:       data.Image,
	}

	if err := postReviewUsecase.postReviewStore.Create(ctx, &models); err != nil {
		return err
	}

	return nil
}
