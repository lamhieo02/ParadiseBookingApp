package requestguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

type RequestGuiderSto interface {
	Create(ctx context.Context, data *entities.RequestGuider) error
	GetByUserID(ctx context.Context, userID int) (*entities.RequestGuider, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *requestguideriomodel.Filter) ([]*entities.RequestGuider, error)
	UpdateByID(ctx context.Context, id int, data *entities.RequestGuider) error
	UpdateWithMap(ctx context.Context, data *entities.RequestGuider, props map[string]interface{}) error
	GetByID(ctx context.Context, id int) (*entities.RequestGuider, error)
}

type AccountSto interface {
	UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error
	GetProfileByID(ctx context.Context, id int) (*entities.Account, error)
}

type requestGuiderUC struct {
	requestGuiderSto RequestGuiderSto
	accountSto       AccountSto
}

func NewRequestGuiderUC(requestGuiderSto RequestGuiderSto, accountSto AccountSto) *requestGuiderUC {
	return &requestGuiderUC{requestGuiderSto, accountSto}
}
