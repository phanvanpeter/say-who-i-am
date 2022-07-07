package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"github.com/phanvanpeter/say-who-i-am/people-api/handlers/peoplehandlers"
	"net/http"
)

type Route struct {
	router *mux.Router
	people *peoplehandlers.People
}

// NewRoutes manages the HTTP requests of the application
func NewRoutes(logger hclog.Logger, validate *data.Validation) *mux.Router {
	r := mux.NewRouter()

	people := peoplehandlers.NewPeople(logger, validate)

	route := Route{
		router: r,
		people: people,
	}

	route.getRoute()
	route.postRoute()
	route.putRoute()
	route.deleteRoute()

	return r
}

func (r *Route) getRoute() {
	getRoute := r.router.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/people", r.people.GetAll)
	getRoute.HandleFunc("/people/{id:[0-9]+}", r.people.Get)
}

func (r *Route) postRoute() {
	postRoute := r.router.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/people", r.people.Create)
	postRoute.Use(r.people.ValidatePerson)
}

func (r *Route) putRoute() {
	putRoute := r.router.Methods(http.MethodPut).Subrouter()
	putRoute.HandleFunc("/people", r.people.Update)
	putRoute.Use(r.people.ValidatePerson)
}

func (r *Route) deleteRoute() {
	deleteRoute := r.router.Methods(http.MethodDelete).Subrouter()
	deleteRoute.HandleFunc("/people/{id:[0-9]+}", r.people.Delete)
}
