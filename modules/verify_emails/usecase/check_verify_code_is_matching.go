package verifyemailsusecase

import (
	"context"
	"paradise-booking/common"

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

	return nil
}
