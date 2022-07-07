package routes

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-images/files"
	"github.com/phanvanpeter/say-who-i-am/people-images/handlers"
	"net/http"
)

func NewRoutes(logger hclog.Logger, storage files.Storage, basePath string) *mux.Router {
	r := mux.NewRouter()

	f := handlers.NewFiles(logger, storage)

	imgPost := r.Methods(http.MethodPost).Subrouter()
	imgPost.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", f.Save)

	imgGet := r.Methods(http.MethodGet).Subrouter()
	imgGet.Handle("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}",
		http.StripPrefix("/images/", http.FileServer(http.Dir(basePath))))

	return r
}
