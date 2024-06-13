package service

import (
	"context"
	"errors"

	"github.com/yuki-maruyama/learn-gh-action/model"
	"github.com/yuki-maruyama/learn-gh-action/util"
)

type ImageService interface {
	GetResizedImageService(ctx context.Context, req model.ImageRequest) (*model.ImageResponse, error)
}

type imageService struct {
}

var _ ImageService = (*imageService)(nil)

func NewImageService() *imageService {
	return &imageService{}
}

func (s *imageService) GetResizedImageService(ctx context.Context, req model.ImageRequest) (*model.ImageResponse, error) {
	img, err := util.LoadImage("./blob/img.jpg")
	if err != nil {
		return nil, errors.New("image load error")
	}
	var size int = 500
	if req.Size != nil {
		size = *req.Size
	}
	resizedImg := util.ResizeImageKeepAspect(img, size)
	res := &model.ImageResponse{
		Image: resizedImg,
	}
	return res, nil
}
