package commentusecase

import (
	"context"
	"errors"
	"paradise-booking/entities"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCommentByPostReviewID(t *testing.T) {
	// init mock
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	mockCommentSto := NewMockCommentStorage(ctrl)
	replyCommentSto := NewMockReplyCommentStorage(ctrl)
	accountSto := NewMockAccountStorage(ctrl)
	// init usecase
	uc := NewCommentUseCase(mockCommentSto, replyCommentSto, accountSto)

	id := gofakeit.Number(1, 100)

	Convey("Test Get Comment By Post Review ID", t, func() {
		Convey("Get comment by post review ID fail", func() {
			mockCommentSto.EXPECT().GetByPostReviewID(ctx, id).Return(nil, errors.New("error"))
			_, err := uc.GetCommentByPostReviewID(ctx, id)
			So(err, ShouldNotBeNil)
		})
		Convey("Get comment by post review ID success", func() {
			comments := []*entities.Comment{}
			model := entities.Comment{}
			if err := gofakeit.Struct(&model); err != nil {
				t.Error("Error to create data test")
			}
			comments = append(comments, &model)
			mockCommentSto.EXPECT().GetByPostReviewID(ctx, id).Return(comments, nil)
			Convey("Get reply comment by comment ID fail", func() {
				replyCommentSto.EXPECT().GetBySourceCommentID(ctx, gomock.Any()).Return(nil, errors.New("error"))
				_, err := uc.GetCommentByPostReviewID(ctx, id)
				So(err, ShouldNotBeNil)
			})
			Convey("Get reply comment by comment ID success", func() {
				replyCommentSto.EXPECT().GetBySourceCommentID(ctx, gomock.Any()).Return([]entities.ReplyComment{}, nil)
				Convey("Get account by ID fail", func() {
					accountSto.EXPECT().GetProfileByID(ctx, gomock.Any()).Return(nil, errors.New("error"))
					_, err := uc.GetCommentByPostReviewID(ctx, id)
					So(err, ShouldNotBeNil)
				})
				Convey("Get account by ID success", func() {
					accountSto.EXPECT().GetProfileByID(ctx, gomock.Any()).Return(&entities.Account{}, nil)
					_, err := uc.GetCommentByPostReviewID(ctx, id)
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
