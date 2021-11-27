package server

import (
	"net/http"

	"github.com/traperwaze/tinywebserver/routes"
)

type Server struct {
	Host string
}

func (server *Server) Start() {
	// * Handle static assets
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fileServer))

	// * register routes
	routes.Register()

	http.ListenAndServe(server.Host, routes.Router)
}
