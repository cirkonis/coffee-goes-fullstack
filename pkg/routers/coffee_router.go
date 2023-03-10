package routers

import (
	"coffee-goes-fullstack/pkg/handlers"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/machine", handlers.GetCoffeeMachine).Methods("GET")
	router.HandleFunc("/machine", handlers.UpdateCoffeeMachine).Methods("PUT")
	router.HandleFunc("/coffees", handlers.GetCoffees).Methods("GET")
	router.HandleFunc("/coffees/{id}", handlers.GetCoffeeByID).Methods("GET")
	return router
}
