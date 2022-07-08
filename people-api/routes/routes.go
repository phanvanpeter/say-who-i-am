package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"github.com/phanvanpeter/say-who-i-am/people-api/handlers/peoplehandlers"
	"github.com/phanvanpeter/say-who-i-am/people-api/handlers/reviewshandlers"
)

type Route struct {
	router  *mux.Router
	people  *peoplehandlers.People
	reviews *reviewshandlers.Reviews
}

// NewRoutes manages the HTTP requests of the application
func NewRoutes(logger hclog.Logger, validate *data.Validation) *mux.Router {
	r := mux.NewRouter()

	route := initRoute(r, logger, validate)

	route.personRoute()
	route.reviewsRoute()

	return r
}

func initRoute(router *mux.Router, logger hclog.Logger, validate *data.Validation) *Route {
	people := peoplehandlers.NewPeople(logger, validate)
	reviews := reviewshandlers.NewReviews(logger)

	return &Route{
		router:  router,
		people:  people,
		reviews: reviews,
	}
}
