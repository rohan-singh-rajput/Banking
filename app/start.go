package app

import (
	"net/http"

	"github.com/rohan-singh-rajput/banking/handlers"
)

func Start() {
	//define routes
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/customers", handlers.GetAllCustomers)

	// starting server
	http.ListenAndServe("localhost:8000", nil)
}
