package requestvendorhandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/entities"
	requestvendoriomodel "paradise-booking/modules/request_vendor/iomodel"
)

type requestVendorUC interface {
	UpsertRequestVendor(ctx context.Context, data *entities.RequestVendor) error
	GetByUserID(ctx context.Context, userID int) (*requestvendoriomodel.GetRequestVendorResp, error)
	ListByFilter(ctx context.Context, paging *common.Paging, filter *requestvendoriomodel.Filter) ([]*requestvendoriomodel.GetRequestVendorResp, error)
	ConfirmRequestVendor(ctx context.Context, requestVendorID int, typeConfirm int) error
}

type requestVendorHandler struct {
	requestVendorUC requestVendorUC
}

func NewRequestVendorHandler(requestVendorUC requestVendorUC) *requestVendorHandler {
	return &requestVendorHandler{requestVendorUC}
}
