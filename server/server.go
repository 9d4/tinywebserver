package server

import (
	"net/http"

	"github.com/traperwaze/tinywebserver/router"
)

type Server struct {
	Host string
}

func (server *Server) Start() {
	// * Handle static assets
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/", http.StripPrefix("/", fileServer))


	router.Init()
	
	http.ListenAndServe(server.Host, nil)
}