package requestvendorusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

type RequestVendorSto interface {
	Create(ctx context.Context, data *entities.RequestVendor) error
	GetByUserID(ctx context.Context, userID int) (*entities.RequestVendor, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *requestvendoriomodel.Filter) ([]*entities.RequestVendor, error)
	UpdateByID(ctx context.Context, id int, data *entities.RequestVendor) error
	UpdateWithMap(ctx context.Context, data *entities.RequestVendor, props map[string]interface{}) error
	GetByID(ctx context.Context, id int) (*entities.RequestVendor, error)
}

type AccountSto interface {
	UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type requestVendorUC struct {
	requestVendorSto RequestVendorSto
	accountSto       AccountSto
}

func NewRequestVendorUC(requestVendorSto RequestVendorSto, accountSto AccountSto) *requestVendorUC {
	return &requestVendorUC{requestVendorSto, accountSto}
}
