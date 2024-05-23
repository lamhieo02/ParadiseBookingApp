package requestguiderusecase

import (
	"context"
	requestguiderconvert "paradise-booking/modules/request_guider/convert"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

func (uc *requestGuiderUC) GetByUserID(ctx context.Context, userID int) ([]*requestguideriomodel.GetRequestGuiderResp, error) {
	data, err := uc.requestGuiderSto.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var res []*requestguideriomodel.GetRequestGuiderResp
	for _, val := range data {
		res = append(res, requestguiderconvert.ConvertRequestGuiderEntityToModel(&val))
	}
	return res, nil
}
