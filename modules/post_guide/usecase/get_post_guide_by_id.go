package postguideusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *postGuideUsecase) GetPostGuideByID(ctx context.Context, id int) (*entities.PostGuide, error) {
	postGuide, err := uc.postGuideCache.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return postGuide, nil
}
