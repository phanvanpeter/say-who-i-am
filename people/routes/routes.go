package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people/handlers"
)

func NewRoutes(logger hclog.Logger) *mux.Router {
	r := mux.NewRouter()

	people := handlers.NewPeople(logger)

	r.HandleFunc("/people", people.GetPeople)

	return r
}
