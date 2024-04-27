package postguideusecase

import (
	"context"
	"paradise-booking/entities"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func (uc *postGuideUsecase) CreatePostGuide(ctx context.Context, data *postguideiomodel.CreatePostGuideReq) error {
	entity := &entities.PostGuide{
		PostOwnerId: data.PostOwnerID,
		TopicID:     data.TopicID,
		Title:       data.Title,
		Description: data.Description,
		Cover:       data.Cover,
		Lat:         data.Lat,
		Lng:         data.Lng,
	}

	if err := uc.postGuideSto.Create(ctx, entity); err != nil {
		return err
	}

	return nil
}
