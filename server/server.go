package server

import (
	"net/http"

	"github.com/traperwaze/tinywebserver/router"
)

type Server struct {
	Host string
}

func (server *Server) Start() {
	router.Init()

	http.ListenAndServe(server.Host, nil)
}