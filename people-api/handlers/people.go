package handlers

import (
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

type People struct {
	logger hclog.Logger
}

func NewPeople(l hclog.Logger) *People {
	return &People{
		logger: l,
	}
}

func (p *People) GetPeople(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] get all people")
	w.Header().Add("Content-Type", "application/json")

	people := data.GetPeople()
	err := data.ToJSON(people, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.logger.Error("Error serializing people data to JSON and sending them to the client", err)
	}
}
