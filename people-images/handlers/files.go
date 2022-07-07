package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-images/files"
	"net/http"
	"path/filepath"
)

type Files struct {
	logger hclog.Logger
	store  files.Storage
}

func NewFiles(logger hclog.Logger, storage files.Storage) *Files {
	return &Files{
		logger: logger,
		store:  storage,
	}
}

func (f *Files) Save(w http.ResponseWriter, r *http.Request) {
	f.logger.Info("Save file")
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.logger.Info("Handle POST", id, "filename", fn)

	f.saveFile(id, fn, w, r)
}

func (f *Files) saveFile(id, fileName string, w http.ResponseWriter, r *http.Request) {
	f.logger.Info("Saving File for ID", id, "filename", fileName)

	fp := filepath.Join(id, fileName)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.logger.Error("Unable to save file", err)

		http.Error(w, fmt.Sprintf("Unable to save file: %s", err), http.StatusInternalServerError)
	}
}
