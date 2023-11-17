package convert

import (
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
)

func ConvertAccountEntityToInfoResp(account *entities.Account) *iomodel.AccountInfoResp {
	return &iomodel.AccountInfoResp{
		Id: account.Id,
		Email: account.Email,
		Username: account.Username,
		FullName: account.FullName,
		Address: account.Address,
		Phone: account.Phone,
		Dob: account.Dob,
	}
}