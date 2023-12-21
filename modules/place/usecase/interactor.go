package placeusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/config"
	"paradise-booking/entities"
	"paradise-booking/modules/place/iomodel"
	googlemapprovider "paradise-booking/provider/googlemap"
)

type PlaceStorage interface {
	Create(ctx context.Context, data *entities.Place) (err error)
	DeleteByID(ctx context.Context, id int) error
	GetPlaceByID(ctx context.Context, id int) (*entities.Place, error)
	ListPlaceByVendorID(ctx context.Context, vendorID int, paging *common.Paging) ([]entities.Place, error)
	ListPlaces(ctx context.Context, paging *common.Paging, filter *iomodel.Filter, address *googlemapprovider.GoogleMapAddress) ([]entities.Place, error)
	UpdateByID(ctx context.Context, id int, data *entities.Place) error
	GetPlaceByCondition(ctx context.Context, condition map[string]interface{}) ([]entities.Place, error)
}

type AccountStorage interface {
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error)
}

type placeUseCase struct {
	placeStorage PlaceStorage
	accountSto   AccountStorage
	cfg          *config.Config
	googleMap    *googlemapprovider.GoogleMap
}

func NewPlaceUseCase(cfg *config.Config, placeSto PlaceStorage, accoutSto AccountStorage, googleMap *googlemapprovider.GoogleMap) *placeUseCase {
	return &placeUseCase{placeSto, accoutSto, cfg, googleMap}
}
