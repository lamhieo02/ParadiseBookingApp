package postguideusecase

import (
	"context"
	postguideconvert "paradise-booking/modules/post_guide/convert"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func (uc *postGuideUsecase) GetPostGuideByID(ctx context.Context, id int) (*postguideiomodel.GetPostGuideResp, error) {
	postGuide, err := uc.postGuideCache.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	owner, err := uc.accountCache.GetProfileByID(ctx, postGuide.PostOwnerId)
	if err != nil {
		return nil, err
	}

	res := postguideconvert.ConvertPostGuideEntityToModel(postGuide, owner)
	// get rating average
	ratingAverage, err := uc.postGuideSto.GetRatingAverageByPostGuideId(ctx, int64(id))
	if err != nil {
		return nil, err
	}

	if ratingAverage == nil {
		defaulRating := 0.0
		ratingAverage = &defaulRating
	}
	res.RatingAverage = *ratingAverage

	// get place related to post guide
	condition := map[string]interface{}{
		"state": postGuide.State,
	}
	placeIds, err := uc.placeSto.ListPlaceIdsByCondition(ctx, 10, condition)
	if err != nil {
		return nil, err
	}

	for _, placeID := range placeIds {
		place, err := uc.placeCache.GetPlaceByID(ctx, placeID)
		if err != nil {
			return nil, err
		}

		res.PlaceRelates = append(res.PlaceRelates, postguideiomodel.PlaceRelate{
			ID:            place.Id,
			Name:          place.Name,
			Description:   place.Description,
			PricePerNight: place.PricePerNight,
			Address:       place.Address,
			Cover:         place.Cover,
			Country:       place.Country,
			State:         place.State,
			District:      place.District,
		})
	}

	return &res, nil
}
