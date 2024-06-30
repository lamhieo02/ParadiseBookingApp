package reporthandler

import (
	"context"
	"paradise-booking/common"
	reportiomodel "paradise-booking/modules/report/iomodel"
)

type reportUseCase interface {
	CreateReport(ctx context.Context, data *reportiomodel.CreateReportReq) error
	GetReportByID(ctx context.Context, id int) (*reportiomodel.GetReportResp, error)
	UpdateReportByID(ctx context.Context, id int, data *reportiomodel.UpdateReportReq) error
	ListReport(ctx context.Context, paging *common.Paging, filter *reportiomodel.Filter) ([]*reportiomodel.GetReportResp, error)
	GetStatisticsPlace(ctx context.Context, req *reportiomodel.GetStatisticPlaceReq, vendorID int) (*reportiomodel.StatisticPlaceResp, error)
	GetStatisticsPostGuide(ctx context.Context, req *reportiomodel.GetStatisticPostGuideReq, postOwnerID int) (*reportiomodel.StatisticPostGuideResp, error)
}

type reportHandler struct {
	reportUseCase reportUseCase
}

func NewReportHandler(reportUseCase reportUseCase) *reportHandler {
	return &reportHandler{reportUseCase: reportUseCase}
}
