package requestguiderusecase

import (
	"context"
	"paradise-booking/common"
	requestguiderconvert "paradise-booking/modules/request_guider/convert"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

func (uc *requestGuiderUC) ListByFilter(ctx context.Context, paging *common.Paging, filter *requestguideriomodel.Filter) ([]*requestguideriomodel.GetRequestGuiderResp, error) {
	paging.Process()
	data, err := uc.requestGuiderSto.ListByFilter(ctx, paging, filter)
	if err != nil {
		return nil, err
	}

	var res []*requestguideriomodel.GetRequestGuiderResp
	for _, val := range data {
		res = append(res, requestguiderconvert.ConvertRequestGuiderEntityToModel(val))
	}

	return res, nil
}
