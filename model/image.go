package model

import "image"

type ImageRequest struct {
	Size *int
}

type ImageResponse struct {
	Image image.Image
}
