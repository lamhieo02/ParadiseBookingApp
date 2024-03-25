package mediahandler

import (
	"context"
	"mime/multipart"
	"paradise-booking/common"
	"paradise-booking/config"
)

type mediaUseCase interface {
	UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*common.Image, error)
}

type mediaHandler struct {
	mediaUC mediaUseCase
	cfg     *config.Config
}

func NewMediaHandler(cfg *config.Config, mediaUseCase mediaUseCase) *mediaHandler {
	return &mediaHandler{mediaUC: mediaUseCase, cfg: cfg}
}
