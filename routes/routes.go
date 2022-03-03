package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/9d4/tinywebserver/config"
	"github.com/9d4/tinywebserver/controllers/index"
	"github.com/9d4/tinywebserver/helper"
	"github.com/9d4/tinywebserver/middlewares/auth"
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
	// Router.HandleFunc("/", index.Store).Methods("POST")
	// Router.HandleFunc("/", index.Index)
	// Router.HandleFunc("/delete/{id}", index.Delete).Methods("POST")
	// Router.HandleFunc("/done/{id}", index.Done)
	// Router.HandleFunc("/undone/{id}", index.Undone)

	auths := Router.PathPrefix("/user").Subrouter()
	auths.Use(authentication_middleware.Auth)

	auths.HandleFunc("/Ampas", index.Index)
	
	auths_guest := Router.PathPrefix("/user").Subrouter()
	auths_guest.Use(authentication_middleware.Guest)
	
	auths_guest.HandleFunc("/login", index.LoginFake)


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
