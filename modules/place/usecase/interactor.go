package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
)

type PlaceStorage interface {
	Create(ctx context.Context, data *entities.Place) (err error)
	DeleteByID(ctx context.Context, id int) error
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceByVendorID(ctx context.Context, vendorID int) ([]entities.Place, error)
	ListPlaces(ctx context.Context, paging *common.Paging, filter *iomodel.Filter) ([]entities.Place, error)
	UpdateByID(ctx context.Context, id int, data *entities.Place) error
}

type AccountStorage interface {
	GetProfile(ctx context.Context, id int) (*entities.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type placeUseCase struct {
	placeStorage PlaceStorage
	accountSto   AccountStorage
	cfg          *config.Config
}

func NewPlaceUseCase(cfg *config.Config, placeSto PlaceStorage, accoutSto AccountStorage) *placeUseCase {
	return &placeUseCase{placeSto, accoutSto, cfg}
}
