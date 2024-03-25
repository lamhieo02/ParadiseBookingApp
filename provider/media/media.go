package mediaprovider

import (
	"context"
	"os"
	"paradise-booking/common"
	"paradise-booking/config"
)

type MediaProvider struct {
}

func NewMediaProvider(config *config.Config) *MediaProvider {
	return &MediaProvider{}
}

func (m *MediaProvider) SaveImage(ctx context.Context, image []byte, dst string) (*common.Image, error) {
	// Create a new file to save the image
	if err := os.WriteFile(dst, image, 0644); err != nil {
		return nil, err
	}
	return nil, nil
}
