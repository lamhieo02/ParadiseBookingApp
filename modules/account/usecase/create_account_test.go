package accountusecase

import (
	"context"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

// func (uc *accountUseCase) CreateAccount(ctx context.Context, accountModel *iomodel.AccountRegister) (result *string, err error) {

// 	// check if email is existed in db
// 	ac, err := uc.accountStorage.GetAccountByEmail(ctx, accountModel.Email)
// 	if err != nil {
// 		if err != gorm.ErrRecordNotFound {
// 			return nil, err
// 		}
// 	}

// 	if ac != nil {
// 		return nil, errors.New("email is existed")
// 	}
// 	// convert from iomodel to entity
// 	accountEntity := convert.ConvertAccountRegisModelToEntity(accountModel)
// 	accountEntity.Status = int(constant.StatusActive)
// 	// hash password before store in db
// 	hashedPassword, err := utils.HashPassword(accountEntity.Password)
// 	if err != nil {
// 		return nil, common.ErrInternal(err)
// 	}

// 	// default in first register account will have role user
// 	accountEntity.Role = int(constant.UserRole)
// 	accountEntity.Password = hashedPassword

// 	paramCreateTx := accountstorage.CreateUserTxParam{
// 		Data: &accountEntity,
// 		AfterCreate: func(data *entities.Account) error {
// 			// after create account success, we will send email to user to verify account
// 			taskPayload := worker.PayloadSendVerifyEmail{
// 				Email: accountEntity.Email,
// 			}
// 			opts := []asynq.Option{
// 				asynq.MaxRetry(10),
// 				asynq.ProcessIn(10 * time.Second),
// 				asynq.Queue(worker.QueueSendVerifyEmail),
// 			}

// 			return uc.taskDistributor.DistributeTaskSendVerifyEmail(ctx, &taskPayload, opts...)
// 		},
// 	}

// 	if err = uc.accountStorage.CreateTx(ctx, paramCreateTx); err != nil {
// 		return nil, err
// 	}

// 	return &accountEntity.Email, nil
// }

func TestCreateAccount(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)
	mockTaskDistributor := NewMockVerifyEmailsUseCase(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, mockTaskDistributor, nil, nil)

	// prepare data to test
	dataTests := make([]iomodel.AccountRegister, 100)
	for i := 0; i < 100; i++ {
		dataTests[i] = iomodel.AccountRegister{
			Email:    gofakeit.Email(),
			Password: gofakeit.Password(true, true, true, true, false, 10),
		}
	}

	Convey("Test Create Account", t, func() {
		for _, tc := range dataTests {
			Convey("Check email is existed", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(&entities.Account{}, nil)

				result, err := uc.CreateAccount(ctx, &tc)
				convey.So(result, convey.ShouldBeNil)
				convey.So(err, convey.ShouldNotBeNil)
				convey.So(err.Error(), convey.ShouldEqual, "email is existed")
			})
			Convey("Create account successfully", func() {
				mockAccountStorage.EXPECT().GetAccountByEmail(ctx, tc.Email).Return(nil, gorm.ErrRecordNotFound)
				mockAccountStorage.EXPECT().CreateTx(ctx, gomock.Any()).Return(nil)
				result, err := uc.CreateAccount(ctx, &tc)
				convey.So(result, convey.ShouldNotBeNil)
				convey.So(err, convey.ShouldBeNil)
			})
		}
	})
}
