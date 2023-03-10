package handlers

import (
	"coffee-goes-fullstack/pkg/data"
	"coffee-goes-fullstack/pkg/models"
	"encoding/json"
	"net/http"
)

func GetCoffeeMachine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.CoffeeMachine)
}

func UpdateCoffeeMachine(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cm models.CoffeeMachine
	err := json.NewDecoder(r.Body).Decode(&cm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.CoffeeMachine = cm

	// Return a JSON response with the updated coffee machine data and a 200 status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.CoffeeMachine)
}
