package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/modules/account/convert"
	"paradise-booking/modules/account/iomodel"
	"paradise-booking/utils"
)

func (uc *accountUseCase) CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error) {

	// convert from iomodel to entity
	accountEntity := convert.ConvertAccountRegisModelToEntity(accountModel)

	// hash password before store in db
	hashedPassword, err := utils.HashPassword(accountEntity.Password)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// default in first register account will have role user
	accountEntity.Role = int(constant.UserRole)

	accountEntity.Password = hashedPassword
	if err = uc.accountStorage.Create(ctx, &accountEntity); err != nil {
		return nil, err
	}
	return &accountEntity.Email, nil
}
