package requestvendorusecase

import (
	"context"
	"paradise-booking/constant"
	requestvendorconvert "paradise-booking/modules/request_vendor/convert"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

func (uc *requestVendorUC) GetByUserID(ctx context.Context, userID int) (*requestvendoriomodel.GetRequestVendorResp, error) {
	data, err := uc.requestVendorSto.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := requestvendorconvert.ConvertRequestVendorEntityToModel(data)
	user, err := uc.accountSto.GetProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res.User.ID = user.Id
	res.User.Role = constant.MapRole[constant.Role(user.Role)]
	res.User.Username = user.Username
	res.User.FullName = user.FullName
	res.User.Email = user.Email
	res.User.Phone = user.Phone
	res.User.Address = user.Address
	res.User.DOB = user.Dob
	res.User.Avt = user.Avatar
	res.User.Bio = user.Bio
	res.User.Created = user.CreatedAt
	res.User.Updated = user.UpdatedAt

	return res, nil

}
