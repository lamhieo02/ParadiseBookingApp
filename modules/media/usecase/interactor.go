package mediausecase

import (
	"paradise-booking/config"
	mediaprovider "paradise-booking/provider/media"
)

type mediaUseCase struct {
	cfg           *config.Config
	mediaProvider *mediaprovider.MediaProvider
}

func NewMediaUseCase(cfg *config.Config, mediaProvider *mediaprovider.MediaProvider) *mediaUseCase {
	return &mediaUseCase{cfg, mediaProvider}
}
