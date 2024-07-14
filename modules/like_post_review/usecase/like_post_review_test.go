package likepostreviewusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	likepostreviewiomodel "paradise-booking/modules/like_post_review/iomodel"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLikePostReview(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockLikePostReviewSto := NewMockLikePostReviewStorage(ctrl)
	// init usecase
	uc := NewLikePostReviewUseCase(mockLikePostReviewSto)

	req := likepostreviewiomodel.LikePostReviewReq{}
	if err := gofakeit.Struct(&req); err != nil {
		t.Error("Error to create data test")
	}

	resp := make([]*entities.LikePostReview, 1)
	model := entities.LikePostReview{}
	if err := gofakeit.Struct(&model); err != nil {
		t.Error("Error to create data test")
	}
	resp[0] = &model
	Convey("Test Like Post Review", t, func() {
		Convey("FindDataByCondition fail", func() {
			mockLikePostReviewSto.EXPECT().FindDataByCondition(ctx, gomock.Any()).Return(nil, errors.New("error1"))
			err := uc.LikePostReview(ctx, &req)
			So(err, ShouldNotBeNil)
		})
		Convey("FindDataByCondition success", func() {
			Convey("len = 0", func() {
				mockLikePostReviewSto.EXPECT().FindDataByCondition(ctx, gomock.Any()).Return([]*entities.LikePostReview{}, nil)
				mockLikePostReviewSto.EXPECT().Create(ctx, gomock.Any()).Return(nil)
				err := uc.LikePostReview(ctx, &req)
				So(err, ShouldBeNil)
			})
			Convey("len > 0", func() {
				mockLikePostReviewSto.EXPECT().FindDataByCondition(ctx, gomock.Any()).Return(resp, nil)
				mockLikePostReviewSto.EXPECT().UpdateWithMap(ctx, gomock.Any(), gomock.Any()).Return(nil)
				err := uc.LikePostReview(ctx, &req)
				So(err, ShouldBeNil)
			})
		})
	})
}
