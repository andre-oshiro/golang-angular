package main

import (
	"net/http"

	"./api"
)

func main() {

	mux := http.NewServeMux()

	// Setup API's
	api.SetupUserAPI(mux)
	api.SetupItemAPI(mux)

	mux.Handle("/", http.FileServer(http.Dir("SPA")))

	http.ListenAndServe(":8080", mux)

}
