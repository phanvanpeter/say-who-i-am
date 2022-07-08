package routes

import "net/http"

func (r *Route) personRoute() {
	getRoute := r.router.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/people", r.people.GetAll)
	getRoute.HandleFunc("/people/{id:[0-9]+}", r.people.GetByID)

	postRoute := r.router.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/people", r.people.Create)
	postRoute.Use(r.people.ValidatePerson)

	putRoute := r.router.Methods(http.MethodPut).Subrouter()
	putRoute.HandleFunc("/people", r.people.Update)
	putRoute.Use(r.people.ValidatePerson)

	deleteRoute := r.router.Methods(http.MethodDelete).Subrouter()
	deleteRoute.HandleFunc("/people/{id:[0-9]+}", r.people.Delete)
}
