package handlers

import (
	"fmt"
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

func (p *People) Create(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] create a new person")

	person := &data.Person{}
	err := data.FromJSON(person, r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid data of a person: %s", err), http.StatusBadRequest)
		return
	}
	data.AddPerson(person)

}
