package reportusecase

import (
	"context"
	"paradise-booking/constant"
	reportconvert "paradise-booking/modules/report/convert"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

func (uc *reportUseCase) GetReportByID(ctx context.Context, id int) (*reportiomodel.GetReportResp, error) {
	report, err := uc.reportSto.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	account, err := uc.accountSto.GetProfileByID(ctx, report.UserID)
	if err != nil {
		return nil, err
	}

	result := reportconvert.ReportEntityToModel(report)
	result.User.ID = account.Id
	result.User.Role = constant.MapRole[constant.Role(account.Role)]
	result.User.Username = account.Username
	result.User.FullName = account.FullName
	result.User.Email = account.Email
	result.User.Phone = account.Phone
	result.User.Address = account.Address
	result.User.DOB = account.Dob
	result.User.Avt = account.Avatar
	result.User.Bio = account.Bio
	result.User.Created = account.CreatedAt
	result.User.Updated = account.UpdatedAt

	return result, nil
}
