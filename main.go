package main

import (
	"github.com/traperwaze/tinywebserver/config"
	"github.com/traperwaze/tinywebserver/database"
	"github.com/traperwaze/tinywebserver/server"
)

func main() {
	config.Init()
	database.Init()

	httpServer := server.Server{Host: config.Get("HOST")}
	httpServer.Start()
}
