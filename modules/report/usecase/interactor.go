package reportusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

type reportStorage interface {
	Create(ctx context.Context, data *entities.Report) error
	GetByID(ctx context.Context, id int) (*entities.Report, error)
	UpdateByID(ctx context.Context, id int, data *entities.Report) error
	ListReport(ctx context.Context, paging *common.Paging, filter *reportiomodel.Filter) ([]*entities.Report, error)
}

type placeCache interface {
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
}

type postGuideCache interface {
	GetByID(ctx context.Context, id int) (*entities.PostGuide, error)
}

type accountCache interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type postReviewSto interface {
	GetByID(ctx context.Context, postReviewID int) (*entities.PostReview, error)
}

type commentSto interface {
	GetByID(ctx context.Context, id int) (*entities.Comment, error)
}

type bookingSto interface {
	ListAllBookingWithCondition(ctx context.Context, condition []common.Condition) ([]entities.Booking, error)
}

type placeSto interface {
	ListPlaceIdsByCondition(ctx context.Context, limit int, condition map[string]interface{}) ([]int, error)
}

type bookingDetailCache interface {
	GetByBookingID(ctx context.Context, bookingId int) (res *entities.BookingDetail, err error)
}

type reportUseCase struct {
	reportSto          reportStorage
	accountCache       accountCache
	placeCache         placeCache
	postGuideCache     postGuideCache
	postReviewSto      postReviewSto
	commentSto         commentSto
	bookingSto         bookingSto
	placeSto           placeSto
	bookingDetailCache bookingDetailCache
}

func NewReportUseCase(
	reportSto reportStorage,
	accountCache accountCache,
	placeCache placeCache,
	postGuideCache postGuideCache,
	postReviewSto postReviewSto,
	commentSto commentSto,
	bookingSto bookingSto,
	placeSto placeSto,
	bookingDetailCache bookingDetailCache) *reportUseCase {
	return &reportUseCase{reportSto: reportSto,
		accountCache:       accountCache,
		placeCache:         placeCache,
		postGuideCache:     postGuideCache,
		postReviewSto:      postReviewSto,
		commentSto:         commentSto,
		bookingSto:         bookingSto,
		placeSto:           placeSto,
		bookingDetailCache: bookingDetailCache}
}
