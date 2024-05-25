package requestguiderusecase

import (
	"context"
	requestguiderconvert "paradise-booking/modules/request_guider/convert"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

func (uc *requestGuiderUC) GetByUserID(ctx context.Context, userID int) (*requestguideriomodel.GetRequestGuiderResp, error) {
	data, err := uc.requestGuiderSto.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := requestguiderconvert.ConvertRequestGuiderEntityToModel(data)
	user, err := uc.accountSto.GetProfileByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	res.User.ID = user.Id
	res.User.Username = user.Username
	res.User.FullName = user.FullName
	res.User.Email = user.Email
	res.User.Phone = user.Phone
	res.User.Address = user.Address
	res.User.DOB = user.Dob

	return res, nil

}
