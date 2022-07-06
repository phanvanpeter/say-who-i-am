package handlers

import (
	"github.com/hashicorp/go-hclog"
)

// People is struct for handling requests regarding the people
type People struct {
	logger hclog.Logger
}

type KeyPerson struct{}

// NewPeople initializes a new handler for people
func NewPeople(l hclog.Logger) *People {
	return &People{
		logger: l,
	}
}
