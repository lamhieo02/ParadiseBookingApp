package amenityusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *amenityUseCase) GetAllConfigAmenity(ctx context.Context, typeInt int) (res []entities.ConfigAmenity, err error) {

	res, err = uc.amenitySto.GetAllAmenityConfig(ctx, typeInt)
	if err != nil {
		return nil, err
	}
	return res, nil

}
