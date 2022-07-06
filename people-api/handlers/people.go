package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people/data"
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
