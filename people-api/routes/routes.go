package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people/handlers"
	"net/http"
)

// NewRoutes manages the HTTP requests of the application
func NewRoutes(logger hclog.Logger) *mux.Router {
	r := mux.NewRouter()

	people := handlers.NewPeople(logger)

	getRoute := r.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/people", people.GetAll)

	postRoute := r.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/people", people.Create)
	postRoute.Use(people.ValidatePerson)

	putRoute := r.Methods(http.MethodPut).Subrouter()
	putRoute.HandleFunc("/people", people.Update)
	putRoute.Use(people.ValidatePerson)

	return r
}
