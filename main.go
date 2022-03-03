package main

import (
	"github.com/9d4/tinywebserver/config"
	"github.com/9d4/tinywebserver/database"
	"github.com/9d4/tinywebserver/server"
)

func main() {
	config.Init()
	database.Init()

	httpServer := server.Server{Host: config.Get("HOST")}
	httpServer.Start()
}
