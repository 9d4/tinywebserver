package server

import (
	"net/http"

	"github.com/traperwaze/tinywebserver/routes"
)

type Server struct {
	Host string
}

func (server *Server) Start() {
	// * register statics
	routes.RegisterStatics()

	// * register routes
	routes.Register()

	http.ListenAndServe(server.Host, routes.Router)
}
