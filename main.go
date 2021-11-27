package main

import (
	"github.com/traperwaze/tinywebserver/server"
)

const HOST = ":8000"

func main() {
	httpServer := server.Server{Host: HOST}

	httpServer.Start();
}