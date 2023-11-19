package placehandler

import (
	"context"
	"paradise-booking/common"
	"paradise-booking/modules/place/iomodel"
)

type placeUseCase interface {
	CreatePlace(ctx context.Context, data *iomodel.CreatePlaceReq, emailVendor string) error
	UpdatePlace(ctx context.Context, data *iomodel.UpdatePlaceReq, placeID int, vendorEmail string) error
	GetPlaceByID(ctx context.Context, placeID int) (result *iomodel.GetPlaceResp, err error)
	ListPlaceByVendor(ctx context.Context, vendorEmail string) (result []iomodel.GetPlaceResp, err error)
	ListPlaceByVendorByID(ctx context.Context, vendorID int) (result []iomodel.GetPlaceResp, err error)
	DeletePlaceByID(ctx context.Context, placeID int, vendorEmail string) (err error)
	ListAllPlace(ctx context.Context, paging *common.Paging, filter *iomodel.Filter) (result []iomodel.GetPlaceResp, err error)
}

type placeHandler struct {
	placeUC placeUseCase
}

func NewPlaceHandler(placeUseCase placeUseCase) *placeHandler {
	return &placeHandler{placeUC: placeUseCase}
}
