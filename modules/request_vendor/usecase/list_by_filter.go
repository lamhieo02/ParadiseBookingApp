package requestvendorusecase

import (
	"context"
	"paradise-booking/common"
	requestvendorconvert "paradise-booking/modules/request_vendor/convert"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

func (uc *requestVendorUC) ListByFilter(ctx context.Context, paging *common.Paging, filter *requestvendoriomodel.Filter) ([]*requestvendoriomodel.GetRequestVendorResp, error) {
	paging.Process()
	data, err := uc.requestVendorSto.ListByFilter(ctx, paging, filter)
	if err != nil {
		return nil, err
	}

	var res []*requestvendoriomodel.GetRequestVendorResp
	for _, val := range data {
		res = append(res, requestvendorconvert.ConvertRequestVendorEntityToModel(val))
	}

	return res, nil
}
