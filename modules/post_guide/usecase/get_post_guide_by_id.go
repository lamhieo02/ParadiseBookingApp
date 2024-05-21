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

	return &res, nil
}
