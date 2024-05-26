package amenityusecase

import (
	"context"
	"paradise-booking/config"
	"paradise-booking/entities"
)

type AmenityStorage interface {
	Delete(ctx context.Context, id int) error
	GetAllAmenityConfig(ctx context.Context, typeInt int) ([]entities.ConfigAmenity, error)
	Create(ctx context.Context, data *entities.Amenity) (res *entities.Amenity, err error)
	ListByObjectID(ctx context.Context, objectID int, objectType int) ([]entities.Amenity, error)
	DeleteByCondition(ctx context.Context, condition map[string]any) error
}

type amenityUseCase struct {
	cfg        *config.Config
	amenitySto AmenityStorage
}

func NewAmenityUseCase(amenitySto AmenityStorage, config *config.Config) *amenityUseCase {
	return &amenityUseCase{amenitySto: amenitySto, cfg: config}
}
