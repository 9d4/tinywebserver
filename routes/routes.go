package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/traperwaze/tinywebserver/controllers/index"
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
	Router.HandleFunc("/", index.Index)
	Router.HandleFunc("/index", index.Index)
	Router.HandleFunc("/indexa", index.Index).Methods("POST")
	

	Router.NotFoundHandler = notFoundHandler()
}

func RegisterStatics() {
	fs := http.FileServer(http.Dir("public"))

	Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
