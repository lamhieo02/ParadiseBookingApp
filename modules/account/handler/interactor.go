package accounthandler

import (
	"context"
	"paradise-booking/modules/account/iomodel"
	jwtprovider "paradise-booking/provider/jwt"
)

type accountUseCase interface {
	CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error)
	LoginAccount(ctx context.Context, accountModel *iomodel.AccountLogin) (toke *jwtprovider.Token, err error)
	UpdatePersonalInforAccountById(ctx context.Context, accountModel *iomodel.AccountUpdatePersonalInfo, id int) (err error)
	GetAccountByEmail(ctx context.Context, email string) (account *iomodel.AccountInfoResp, err error)
}

type accountHandler struct {
	accountUC accountUseCase
}

func NewAccountHandler(accountUseCase accountUseCase) *accountHandler {
	return &accountHandler{accountUC: accountUseCase}
}
