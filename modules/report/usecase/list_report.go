package reportusecase

import (
	"context"
	"paradise-booking/common"
	reportconvert "paradise-booking/modules/report/convert"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) ListReport(ctx context.Context, paging *common.Paging, filter *reportiomodel.Filter) ([]*reportiomodel.GetReportResp, error) {
	paging.Process()
	data, err := uc.reportSto.ListReport(ctx, paging, filter)
	if err != nil {
		return nil, err
	}

	var result []*reportiomodel.GetReportResp
	for _, report := range data {
		result = append(result, reportconvert.ReportEntityToModel(report))
	}

	return result, nil
}
