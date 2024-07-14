package amenityusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDeleteAmenityById(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAmenityStorage := NewMockAmenityStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewAmenityUseCase(mockAmenityStorage, &cfg)

	id := gofakeit.Number(1, 100)

	Convey("Test Delete Amenity By ID", t, func() {
		Convey("Delete Amenity By ID Fail", func() {
			mockAmenityStorage.EXPECT().Delete(ctx, id).Return(errors.New("error"))
			err := uc.DeleteAmenityById(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Delete Amenity By ID Success", func() {
			mockAmenityStorage.EXPECT().Delete(ctx, id).Return(nil)
			err := uc.DeleteAmenityById(ctx, id)
			So(err, ShouldBeNil)
		})
	})
}
