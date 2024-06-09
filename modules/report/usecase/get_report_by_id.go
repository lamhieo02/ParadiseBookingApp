package reportusecase

import (
	"context"
	reportconvert "paradise-booking/modules/report/convert"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) GetReportByID(ctx context.Context, id int) (*reportiomodel.GetReportResp, error) {
	report, err := uc.reportSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return reportconvert.ReportEntityToModel(report), nil
}
