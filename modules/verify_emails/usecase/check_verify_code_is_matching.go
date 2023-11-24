package verifyemailsusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"

	"gorm.io/gorm"
)

func (uc *verifyEmailsUseCase) CheckVerifyCodeIsMatching(ctx context.Context, email string, code string) error {
	// check if verify code and email is matching
	data, err := uc.verifyEmailsStore.Get(ctx, email, code)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.ErrVerifyCodeIsNotMatching("verify code", nil)
		}
		return err
	}

	// check if verify code is expired
	if data.IsExpired() {
		return common.ErrExpiredVerifyCode("verify code", err)
	}

	// if all is ok => update status to verified
	account := &entities.Account{
		IsEmailVerified: 1,
	}

	err = uc.accountStore.UpdateIsVerifyEmailByEmail(ctx, account.Email)
	if err != nil {
		return common.ErrCannotUpdateEntity("account", err)
	}

	return nil
}
