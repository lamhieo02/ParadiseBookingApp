package postreviewusecase

import (
	"context"
	"paradise-booking/entities"
	postreviewiomodel "paradise-booking/modules/post_review/iomodel"
	"strings"
)

func (postReviewUsecase *postReviewUsecase) UpdatePostReview(ctx context.Context, data *postreviewiomodel.UpdatePostReviewReq) error {
	// old data to check exist
	_, err := postReviewUsecase.postReviewStore.GetByID(ctx, int(data.PostReviewID))
	if err != nil {
		return err
	}

	// update data
	newData := &entities.PostReview{
		Title:   data.Title,
		Topic:   data.Topic,
		Content: data.Content,
		Image:   strings.Join(data.Images, ","),
		Videos:  strings.Join(data.Videos, ","),
		Lat:     data.Lat,
		Lng:     data.Lng,
	}

	if err := postReviewUsecase.postReviewStore.UpdateByID(ctx, int(data.PostReviewID), newData); err != nil {
		return err
	}

	return nil
}
