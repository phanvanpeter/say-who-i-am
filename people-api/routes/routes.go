package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"github.com/phanvanpeter/say-who-i-am/people-api/handlers"
	"net/http"
)

// NewRoutes manages the HTTP requests of the application
func NewRoutes(logger hclog.Logger, validate *data.Validation) *mux.Router {
	r := mux.NewRouter()

	people := handlers.NewPeople(logger, validate)

	getRoute := r.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/people", people.GetAll)
	getRoute.HandleFunc("/people/{id:[0-9]+}", people.Get)

	postRoute := r.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/people", people.Create)
	postRoute.Use(people.ValidatePerson)

	putRoute := r.Methods(http.MethodPut).Subrouter()
	putRoute.HandleFunc("/people", people.Update)
	putRoute.Use(people.ValidatePerson)

	deleteRoute := r.Methods(http.MethodDelete).Subrouter()
	deleteRoute.HandleFunc("/people/{id:[0-9]+}", people.Delete)

	return r
}
