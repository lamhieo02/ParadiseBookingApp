package postguideusecase

import (
	"context"
	postguideconvert "paradise-booking/modules/post_guide/convert"
	postguideiomodel "paradise-booking/modules/post_guide/iomodel"
)

func (uc *postGuideUsecase) UpdatePostGuideByID(ctx context.Context, id int, postGuide *postguideiomodel.UpdatePostGuideReq) error {
	postGuideEntity := postguideconvert.ConvertPostGuideUpdateToEntity(postGuide)
	if err := uc.postGuideSto.UpdateByID(ctx, id, postGuideEntity); err != nil {
		return err
	}

	return nil
}
