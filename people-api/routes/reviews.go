package routes

import "net/http"

func (r *Route) reviewsRoute() {
	getRoute := r.router.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/reviews", r.reviews.GetAll)
	getRoute.HandleFunc("/reviews/{id:[0-9]+}", r.reviews.GetByID)

	postRoute := r.router.Methods(http.MethodPost).Subrouter()
	postRoute.HandleFunc("/reviews", r.reviews.Create)
	postRoute.Use(r.reviews.ValidateReview)
	//
	//putRoute := r.router.Methods(http.MethodPut).Subrouter()
	//putRoute.HandleFunc("/people", r.people.Update)
	//putRoute.Use(r.people.ValidatePerson)
	//
	//deleteRoute := r.router.Methods(http.MethodDelete).Subrouter()
	//deleteRoute.HandleFunc("/people/{id:[0-9]+}", r.people.Delete)
}
