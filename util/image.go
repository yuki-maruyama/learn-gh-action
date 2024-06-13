package util

import (
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/draw"
)

func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ResizeImageKeepAspect(img image.Image, size int) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	if width > height {
		height = height * size / width
		width = size
	} else {
		width = width * size / height
		height = size
	}

	return ResizeImage(img, width, height)
}

func ResizeImage(img image.Image, width, height int) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}
