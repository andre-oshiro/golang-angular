package api

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

// UserAPI is a API Struct for User
type UserAPI struct{}

// SetupUserAPI prepare API for User
func SetupUserAPI(mux *http.ServeMux) {
	userAPI := UserAPI{}
	mux.HandleFunc("/api/user", userAPI.List)
	mux.HandleFunc("/api/user/{id}", userAPI.Get)
}

// Get function get User
func (api *UserAPI) Get(w http.ResponseWriter, r *http.Request) {
	uStore := model.UserStore{}
	user, err := uStore.Get()
	if err != nil {
		log.Fatal("uStore.Get(). Error: ", err)
	}

	json.NewEncoder(w).Encode(user)
}

// List function get a list of Users
func (api *UserAPI) List(w http.ResponseWriter, r *http.Request) {
	uStore := model.UserStore{}
	users, err := uStore.List()
	if err != nil {
		log.Fatal("uStore.List(). Error: ", err)
	}

	json.NewEncoder(w).Encode(users)
}
