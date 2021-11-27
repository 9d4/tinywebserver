package router

import (
	"net/http"

	indexController "github.com/traperwaze/tinywebserver/controllers/index"
)

func Init() {
	add("/index", indexController.Index)
	add("/index/", indexController.Else)
}

type responder func(w http.ResponseWriter, r *http.Request)
type HandlerFunc responder

func add(path string, handler responder) {
	http.HandleFunc(path, handler)
}
