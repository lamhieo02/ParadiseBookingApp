package placehandler

import (
	"context"
	"paradise-booking/modules/place/iomodel"
)

type placeUseCase interface {
	CreatePlace(ctx context.Context, data *iomodel.CreatePlaceReq, emailVendor string) error
	UpdatePlace(ctx context.Context, data *iomodel.UpdatePlaceReq, placeID int, vendorEmail string) error
}

type placeHandler struct {
	placeUC placeUseCase
}

func NewPlaceHandler(placeUseCase placeUseCase) *placeHandler {
	return &placeHandler{placeUC: placeUseCase}
}
