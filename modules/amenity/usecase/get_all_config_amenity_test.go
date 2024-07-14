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

func TestGetAllConfigAmenity(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAmenityStorage := NewMockAmenityStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewAmenityUseCase(mockAmenityStorage, &cfg)

	typeInt := gofakeit.Number(1, 100)

	Convey("Test Get All Config Amenity", t, func() {
		Convey("Get All Config Amenity Fail", func() {
			mockAmenityStorage.EXPECT().GetAllAmenityConfig(ctx, typeInt).Return(nil, errors.New("error"))
			_, err := uc.GetAllConfigAmenity(ctx, typeInt)
			So(err, ShouldNotBeNil)
		})
		Convey("Get All Config Amenity Success", func() {
			mockAmenityStorage.EXPECT().GetAllAmenityConfig(ctx, typeInt).Return(nil, nil)
			_, err := uc.GetAllConfigAmenity(ctx, typeInt)
			So(err, ShouldBeNil)
		})
	})

}
