package requestguiderhandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestguideriomodel "paradise-booking/modules/request_guider/iomodel"
)

type requestGuiderUC interface {
	UpsertRequestGuider(ctx context.Context, data *entities.RequestGuider) error
	GetByUserID(ctx context.Context, userID int) (*requestguideriomodel.GetRequestGuiderResp, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *requestguideriomodel.Filter) ([]*requestguideriomodel.GetRequestGuiderResp, error)
	ConfirmRequestGuider(ctx context.Context, requestGuiderID int, typeConfirm int) error
}

type RequestGuiderHandler struct {
	requestGuiderUC requestGuiderUC
}

func NewRequestGuiderHandler(requestGuiderUC requestGuiderUC) *RequestGuiderHandler {
	return &RequestGuiderHandler{requestGuiderUC}
}
