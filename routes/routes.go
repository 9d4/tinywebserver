package routes

import (
	"github.com/gorilla/mux"
	"github.com/traperwaze/tinywebserver/controllers/index"
)

// Instance of *mux.Router.
// Use after Register is called
var Router *mux.Router


// Call this function to registers all routes
func Register() {
	Router = mux.NewRouter()
	
	Router.HandleFunc("/", index.Index)
}


