package amenityusecase

import (
	"context"
	"errors"
	"paradise-booking/config"
	"paradise-booking/modules/amenity/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateAmenity(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	mockAmenityStorage := NewMockAmenityStorage(ctrl)

	// init usecase
	cfg := config.Config{}
	uc := NewAmenityUseCase(mockAmenityStorage, &cfg)

	// prepare data to test
	var dataTest iomodel.CreateAmenityReq
	err := gofakeit.Struct(&dataTest)
	if err != nil {
		t.Error("Error when create fake data")
	}

	Convey("Test Create Amenity", t, func() {
		Convey("Create Amenity Fail", func() {
			mockAmenityStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil, errors.New("error12")).Times(1)
			err := uc.CreateAmenity(ctx, &dataTest)
			So(err, ShouldNotBeNil)
		})
		Convey("Create Amenity Success", func() {
			mockAmenityStorage.EXPECT().Create(ctx, gomock.Any()).Return(nil, nil).AnyTimes()
			err := uc.CreateAmenity(ctx, &dataTest)
			So(err, ShouldBeNil)
		})
	})

}
