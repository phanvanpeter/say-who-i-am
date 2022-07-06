package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func NewRoutes(logger hclog.Logger) *mux.Router {
	r := mux.NewRouter()

	return r
}