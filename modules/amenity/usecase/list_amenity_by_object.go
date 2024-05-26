package amenityusecase

import (
	"context"
	"paradise-booking/entities"
)

func (uc *amenityUseCase) ListAmenityByObjectID(ctx context.Context, objectID int, objectType int) (res []entities.Amenity, err error) {

	res, err = uc.amenitySto.ListByObjectID(ctx, objectID, objectType)
	if err != nil {
		return nil, err
	}
	return res, nil

}
