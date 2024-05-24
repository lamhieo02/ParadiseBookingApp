package requestguiderusecase

import (
	"context"
	requestguiderconvert "paradise-booking/modules/request_guider/convert"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

func (uc *requestGuiderUC) GetByUserID(ctx context.Context, userID int) (*requestguideriomodel.GetRequestGuiderResp, error) {
	data, err := uc.requestGuiderSto.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return requestguiderconvert.ConvertRequestGuiderEntityToModel(data), nil
}
