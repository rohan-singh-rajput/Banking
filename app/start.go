package app

import (
	"banking/handlers"
	"net/http"
)

func Start() {
	//define routes
	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/customers", handlers.GetAllCustomers)

	// starting server
	http.ListenAndServe("localhost:8000", nil)
}
