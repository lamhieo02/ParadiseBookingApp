package reportusecase

import (
	"context"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) CreateReport(ctx context.Context, data *reportiomodel.CreateReportReq) error {
	reportToCreate := data.ToEntity()
	err := uc.reportSto.Create(ctx, reportToCreate)
	if err != nil {
		return err
	}

	return nil
}
