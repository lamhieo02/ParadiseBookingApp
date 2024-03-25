package mediausecase

import (
	"context"
	"fmt"
	"mime/multipart"
	"paradise-booking/common"
	"strconv"
	"time"
)

// func (uc *uploadUseCase) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*common.Image, error) {
// 	fileName := fileHeader.Filename

// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		panic(common.ErrBadRequest(err))
// 	}

// 	defer file.Close()

// 	dataBytes := make([]byte, fileHeader.Size)
// 	if _, err := file.Read(dataBytes); err != nil {
// 		panic(common.ErrBadRequest(err))
// 	}

// 	pathFile := fmt.Sprintf("%s/%s", uc.cfg.AWS.S3FolderImages, fileName)
// 	img, err := uc.s3Provider.PutObject(ctx, dataBytes, pathFile)
// 	if err != nil {
// 		panic(common.ErrBadRequest(err))
// 	}

// 	return img, nil
// }

func (uc *mediaUseCase) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*common.Image, error) {
	fileName := strconv.Itoa(time.Now().Nanosecond()) + fileHeader.Filename

	file, err := fileHeader.Open()
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	defer file.Close()

	dataBytes := make([]byte, fileHeader.Size)
	if _, err := file.Read(dataBytes); err != nil {
		panic(common.ErrBadRequest(err))
	}

	pathFile := fmt.Sprintf("./%s/%s", uc.cfg.Image.ImageFolder, fileName)
	_, err = uc.mediaProvider.SaveImage(ctx, dataBytes, pathFile)
	if err != nil {
		panic(common.ErrBadRequest(err))
	}

	return &common.Image{Url: fileName}, nil
}
