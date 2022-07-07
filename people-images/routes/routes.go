package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-images/files"
	"github.com/phanvanpeter/say-who-i-am/people-images/handlers"
	"net/http"
)

func NewRoutes(logger hclog.Logger, storage files.Storage) *mux.Router {
	r := mux.NewRouter()

	f := handlers.NewFiles(logger, storage)

	imageRouter := r.Methods(http.MethodPost).Subrouter()
	imageRouter.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", f.Save)

	return r
}
