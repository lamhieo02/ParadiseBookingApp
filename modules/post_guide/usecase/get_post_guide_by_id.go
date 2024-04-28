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
	return &res, nil
}
