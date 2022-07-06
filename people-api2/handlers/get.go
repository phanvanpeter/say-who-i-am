package handlers

import (
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

// GetAll gets all the people from the storage and pass it to the client
func (p *People) GetAll(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] get all people")
	w.Header().Add("Content-Type", "application/json")

	people := data.GetPeople()
	err := data.ToJSON(people, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.logger.Error("Error serializing people data to JSON and sending them to the client", err)
	}
}
