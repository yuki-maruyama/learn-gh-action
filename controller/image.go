package controller

import (
	"image/jpeg"
	"net/http"

	"github.com/yuki-maruyama/learn-gh-action/model"
	"github.com/yuki-maruyama/learn-gh-action/service"
	"github.com/yuki-maruyama/learn-gh-action/util"
)

type imageController struct {
	service service.ImageService
}

func NewImageController(s service.ImageService) *imageController {
	return &imageController{service: s}
}

func (c *imageController) GetImageHandler(w http.ResponseWriter, r *http.Request) {
	req := model.ImageRequest{
		Size: util.StrToIntPtr(r.FormValue("size")),
	}
	res, err := c.service.GetResizedImageService(r.Context(), req)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.WriteHeader(200)
	err = jpeg.Encode(w, res.Image, &jpeg.Options{})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
