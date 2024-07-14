package accountusecase

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestChangeStatusAccount(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAccountStorage := NewMockAccountStorage(ctrl)

	// init usecase
	uc := NewUserUseCase(nil, mockAccountStorage, nil, nil, nil)

	// prepare data to test
	accountID := gofakeit.Number(1, 100)
	status := gofakeit.Number(0, 1)
	Convey("Test Change Status Account", t, func() {
		Convey("Update status account fail", func() {
			mockAccountStorage.EXPECT().UpdateAccountById(ctx, accountID, gomock.Any()).Return(errors.New("error"))
			err := uc.ChangeStatusAccount(ctx, accountID, status)
			So(err, ShouldNotBeNil)
		})
		Convey("Update status account success", func() {
			mockAccountStorage.EXPECT().UpdateAccountById(ctx, accountID, gomock.Any()).Return(nil)
			err := uc.ChangeStatusAccount(ctx, accountID, status)
			So(err, ShouldBeNil)
		})
	})
}
