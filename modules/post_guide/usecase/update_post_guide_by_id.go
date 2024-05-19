package postguideusecase

import (
	"context"
	postguideconvert "paradise-booking/modules/post_guide/convert"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
	"time"
)

func (uc *postGuideUsecase) UpdatePostGuideByID(ctx context.Context, id int, postGuide *postguideiomodel.UpdatePostGuideReq) error {
	postGuideEntity := postguideconvert.ConvertPostGuideUpdateToEntity(postGuide)
	if err := uc.postGuideSto.UpdateByID(ctx, id, postGuideEntity); err != nil {
		return err
	}

	// update in cache
	postGuideData, err := uc.postGuideSto.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := uc.redisCache.Set(ctx, postGuideData.CacheKey(), &postGuideData, 24*5*time.Hour); err != nil {
		return err
	}

	return nil
}
