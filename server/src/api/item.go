package api

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
)

// ItemAPI is a API Struct for Item
type ItemAPI struct{}

// SetupItemAPI prepare API for Item
func SetupItemAPI(mux *http.ServeMux) {
	itemAPI := ItemAPI{}
	mux.HandleFunc("/api/item", itemAPI.List)
	mux.HandleFunc("/api/item/{id}", itemAPI.Get)
}

// Get function get Item
func (api *ItemAPI) Get(w http.ResponseWriter, r *http.Request) {
	iStore := model.ItemStore{}
	item, err := iStore.Get()
	if err != nil {
		log.Fatal("iStore.Get(). Error: ", err)
	}

	json.NewEncoder(w).Encode(item)
}

// List function gets a List of Items
func (api *ItemAPI) List(w http.ResponseWriter, r *http.Request) {
	iStore := model.ItemStore{}
	items, err := iStore.List()
	if err != nil {
		log.Fatal("iStore.List(). Error: ", err)
	}

	json.NewEncoder(w).Encode(items)
}
