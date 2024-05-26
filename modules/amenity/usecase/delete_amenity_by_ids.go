package amenityusecase

import (
	"context"
	"paradise-booking/modules/amenity/iomodel"
)

func (u *amenityUseCase) DeleteAmenityByListId(ctx context.Context, req *iomodel.DeleteAmenityReq) error {

	for _, id := range req.ListConfigAmenityId {
		condition := map[string]any{
			"object_id":         req.ObjectID,
			"object_type":       req.ObjectType,
			"config_amenity_id": id,
		}
		err := u.amenitySto.DeleteByCondition(ctx, condition)
		if err != nil {
			return err
		}
	}

	return nil
}
