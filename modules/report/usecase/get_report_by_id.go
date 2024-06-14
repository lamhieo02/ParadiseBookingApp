package reportusecase

import (
	"context"
	"paradise-booking/constant"
	"paradise-booking/entities"
	reportconvert "paradise-booking/modules/report/convert"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) GetReportByID(ctx context.Context, id int) (*reportiomodel.GetReportResp, error) {
	report, err := uc.reportSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	account, err := uc.accountCache.GetProfileByID(ctx, report.UserID)
	if err != nil {
		return nil, err
	}

	result := reportconvert.ReportEntityToModel(report)

	uc.getDataUser(result, account)

	return result, nil
}

func (uc *reportUseCase) getDataUser(reportData *reportiomodel.GetReportResp, account *entities.Account) {
	reportData.User.ID = account.Id
	reportData.User.Role = constant.MapRole[constant.Role(account.Role)]
	reportData.User.Username = account.Username
	reportData.User.FullName = account.FullName
	reportData.User.Email = account.Email
	reportData.User.Phone = account.Phone
	reportData.User.Address = account.Address
	reportData.User.DOB = account.Dob
	reportData.User.Avt = account.Avatar
	reportData.User.Bio = account.Bio
	reportData.User.Created = account.CreatedAt
	reportData.User.Updated = account.UpdatedAt
}
