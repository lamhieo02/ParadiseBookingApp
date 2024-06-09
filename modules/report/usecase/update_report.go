package reportusecase

import (
	"context"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) UpdateReportByID(ctx context.Context, id int, data *reportiomodel.UpdateReportReq) error {
	reportToUpdate := data.ToEntity()
	err := uc.reportSto.UpdateByID(ctx, id, reportToUpdate)
	if err != nil {
		return err
	}

	return nil
}
