package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/traperwaze/tinywebserver/config"
	"github.com/traperwaze/tinywebserver/controllers/index"
	"github.com/traperwaze/tinywebserver/helper"
)

// Instance of *mux.Router.
// Use after Register is called
var Router *mux.Router

func init() {
	Router = mux.NewRouter()
}

func notFoundHandler() http.HandlerFunc {
	var f http.HandlerFunc = index.Else
	return f
}

// register routes here
func Register() {
	Router.HandleFunc("/", index.Store).Methods("POST")
	Router.HandleFunc("/", index.Index)
	Router.HandleFunc("/delete/{id}", index.Delete).Methods("POST")
	Router.HandleFunc("/done/{id}", index.Done)
	Router.HandleFunc("/undone/{id}", index.Undone)
	

	Router.NotFoundHandler = notFoundHandler()
}

func RegisterStatics() {
	publicDir := config.Get("PUBLIC_DIR")
	publicPath := "/static"

	if helper.IsEmptyString(publicDir) {
		publicDir = "public"
	}

	fs := http.FileServer(http.Dir(publicDir))
	Router.PathPrefix(publicPath).Handler(http.StripPrefix(publicPath, fs))
}
