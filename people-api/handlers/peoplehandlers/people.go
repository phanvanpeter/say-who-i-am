package peoplehandlers

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
	"strconv"
)

// People is struct for handling requests regarding the people
type People struct {
	logger   hclog.Logger
	validate *data.Validation
}

type KeyPerson struct{}

// NewPeople initializes a new handler for people
func NewPeople(l hclog.Logger, validate *data.Validation) *People {
	return &People{
		logger:   l,
		validate: validate,
	}
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func GetPersonID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}

	return id, nil
}
