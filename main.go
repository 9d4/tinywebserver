package main

import (
	"github.com/traperwaze/tinywebserver/server"
)

const HOST = ":8000"

func main() {
	_server := server.Server{Host: HOST}
	_server.Start();


	// http.ListenAndServe(getNormalizedHost(), nil)
}