package reportusecase

import (
	"context"
	"paradise-booking/entities"
)

type reportStorage interface {
	Create(ctx context.Context, data *entities.Report) error
	GetByID(ctx context.Context, id int) (*entities.Report, error)
	UpdateByID(ctx context.Context, id int, data *entities.Report) error
}

type reportUseCase struct {
	reportSto reportStorage
}

func NewReportUseCase(reportSto reportStorage) *reportUseCase {
	return &reportUseCase{reportSto: reportSto}
}
