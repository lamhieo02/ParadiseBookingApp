package accountusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/account/iomodel"
	"paradise-booking/provider/cache"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUpdatePersonalInforAccountById(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	mockcache := cache.NewMockCache(ctrl)
	// init usecase
	cfg := config.Config{}
	uc := NewUserUseCase(&cfg, mockAccountStorage, nil, nil, mockcache)

	// prepare data to test
	id := gofakeit.Number(1, 100)
	accountModel := &iomodel.AccountUpdatePersonalInfo{}
	accountModel.FullName = gofakeit.Name()
	accountModel.Phone = gofakeit.Phone()
	accountModel.Address = gofakeit.Address().Address
	accountModel.Avt = gofakeit.URL()
	accountModel.Username = gofakeit.Username()

	Convey("Test Update Personal Info Account By ID", t, func() {
		Convey("Update Account By ID Fail", func() {
			mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(errors.New("error"))
			err := uc.UpdatePersonalInforAccountById(ctx, accountModel, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Update Account By ID Success", func() {
			mockAccountStorage.EXPECT().UpdateAccountById(ctx, gomock.Any(), gomock.Any()).Return(nil)
			accountModel2 := &entities.Account{}
			accountModel2.Id = 1
			accountModel2.Email = gofakeit.Email()
			mockAccountStorage.EXPECT().GetProfileByID(ctx, id).Return(accountModel2, nil)
			mockcache.EXPECT().Set(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).MaxTimes(2)
			err := uc.UpdatePersonalInforAccountById(ctx, accountModel, id)
			So(err, ShouldBeNil)
		})
	})

}
