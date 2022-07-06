package handlers

import (
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

// GetAll gets all the people from the storage and passes it to the client
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

// Get gets the person with a given ID from the storage and passes it to the client
func (p *People) Get(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] get all people")
	w.Header().Add("Content-Type", "application/json")

	id, err := GetPersonID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		p.logger.Error("Invalid person ID", err)
		return
	}

	person, err := data.GetPerson(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		p.logger.Error("Person not found", err)
		return
	}
	err = data.ToJSON(person, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		p.logger.Error("Error serializing person data to JSON and sending them to the client", err)
		return
	}
}
