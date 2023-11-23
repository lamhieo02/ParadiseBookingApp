package accountusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/constant"
	"paradise-booking/modules/account/convert"
	"paradise-booking/modules/account/iomodel"
	"paradise-booking/utils"
	"paradise-booking/worker"
	"time"

	"github.com/hibiken/asynq"
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

	//TODO: use db transaction to make sure create account success and send email verify success

	// after create account success, we will send email to user to verify account
	taskPayload := worker.PayloadSendVerifyEmail{
		Email: accountEntity.Email,
	}

	opts := []asynq.Option{
		asynq.MaxRetry(5),
		asynq.ProcessIn(10 * time.Second),
	}

	if err := uc.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &taskPayload, opts...); err != nil {
		return nil, err
	}

	return &accountEntity.Email, nil
}
