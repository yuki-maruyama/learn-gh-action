package controller

import (
	"log"
	"net/http"
)

func errorHandler(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(message))
	if err != nil {
		log.Print(err.Error())
	}
}
