package server

import (
	"fmt"
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

	fmt.Println("Listening in", server.Host)
	http.ListenAndServe(server.Host, routes.Router)
}
