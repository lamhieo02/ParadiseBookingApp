package requestguiderusecase

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

type RequestGuiderSto interface {
	Create(ctx context.Context, data *entities.RequestGuider) error
	GetByUserID(ctx context.Context, userID int) ([]entities.RequestGuider, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *requestguideriomodel.Filter) ([]*entities.RequestGuider, error)
}

type requestGuiderUC struct {
	requestGuiderSto RequestGuiderSto
}

func NewRequestGuiderUC(requestGuiderSto RequestGuiderSto) *requestGuiderUC {
	return &requestGuiderUC{requestGuiderSto}
}
