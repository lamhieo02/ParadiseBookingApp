package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	jwtprovider "paradise-booking/provider/jwt"
	"paradise-booking/utils"

	"gorm.io/gorm"
)

func (uc *accountUseCase) LoginAccount(ctx context.Context, accountModel *iomodel.AccountLogin) (token *jwtprovider.Token, err error) {
	if accountModel.Type == constant.LoginWithGoogle {
		return uc.loginAccountGoogle(ctx, accountModel)
	}
	// find account by email
	account, err := uc.accountStorage.GetAccountByEmail(ctx, accountModel.Email)
	if err != nil {
		return nil, common.ErrEmailNotExist(account.TableName(), err)
	}

	// check status account
	if account.Status != constant.StatusActive {
		return nil, common.ErrAccountIsNotActive(account.TableName(), errors.New("account is not active"))
	}

	// check verify account
	if account.IsEmailVerified == 0 {
		return nil, common.ErrAccountIsNotVerify(account.TableName(), errors.New("account is not verify"))
	}

	// Compare password of user with hashed password in db
	if err := utils.Compare(account.Password, accountModel.Password); err != nil {
		return nil, common.ErrEmailOrPasswordInvalid(account.TableName(), err)
	}

	// generate toke
	token, err = jwtprovider.GenerateJWT(jwtprovider.TokenPayload{Role: account.Role, Email: account.Email}, uc.cfg)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return token, nil
}

func (uc *accountUseCase) loginAccountGoogle(ctx context.Context, accountModel *iomodel.AccountLogin) (token *jwtprovider.Token, err error) {
	// find account by email
	role := int(constant.UserRole)
	account, err := uc.accountStorage.GetAccountByEmail(ctx, accountModel.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// create new account
			newUser := entities.Account{
				Username:        accountModel.FullName,
				FullName:        accountModel.FullName,
				Email:           accountModel.Email,
				Role:            int(constant.UserRole),
				Status:          constant.StatusActive,
				IsEmailVerified: constant.TypeVerifyEmail,
				Avatar:          accountModel.Avatar,
			}
			if err := uc.accountStorage.Create(ctx, &newUser); err != nil {
				return nil, err
			}
		} else {
			return nil, common.ErrEmailNotExist(entities.Account{}.TableName(), err)
		}
	} else {
		role = account.Role
	}

	// generate toke
	token, err = jwtprovider.GenerateJWT(jwtprovider.TokenPayload{Role: role, Email: accountModel.Email}, uc.cfg)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return token, nil
}
