package handlers

import (
	"github.com/hashicorp/go-hclog"
)

type People struct {
	logger hclog.Logger
}

func NewPeople(l hclog.Logger) *People {
	return &People{
		logger: l,
	}
}
