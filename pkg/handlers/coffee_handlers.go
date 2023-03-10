package handlers

import (
	"coffee-goes-fullstack/pkg/data"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetCoffees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Coffees)
}

func GetCoffeeByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	for _, coffee := range data.Coffees {
		if coffee.ID == id {
			json.NewEncoder(w).Encode(coffee)
			return
		}
	}
	http.NotFound(w, r)
}
