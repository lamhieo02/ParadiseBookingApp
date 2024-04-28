package postguideusecase

import (
	"context"
)

func (uc *postGuideUsecase) DeletePostGuideByID(ctx context.Context, id int) error {
	if err := uc.postGuideSto.DeleteByID(ctx, id); err != nil {
		return err
	}

	return nil
}
