package router

import (
	"net/http"

	"github.com/yuki-maruyama/learn-gh-action/api/middleware"
	"github.com/yuki-maruyama/learn-gh-action/controller"
	"github.com/yuki-maruyama/learn-gh-action/service"
)

func NewRouter() http.Handler {
	imageService := service.NewImageService()
	imageController := controller.NewImageController(imageService)

	mux := http.NewServeMux()
	middlewares := []func(http.Handler) http.Handler{
		middleware.LoggingMiddleware,
	}

	handleFunc := func(pattern string, handler func(http.ResponseWriter, *http.Request)) {
		mux.Handle(pattern, http.HandlerFunc(handler))
	}

	handleFunc("GET /", imageController.GetImageHandler)

	return applyMiddleware(mux, middlewares...)
}

func applyMiddleware(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
