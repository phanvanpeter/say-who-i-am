package peoplehandlers

import (
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

// Update updates a person in the storage with a given ID from the put request
func (p *People) Update(w http.ResponseWriter, r *http.Request) {
	person := r.Context().Value(KeyPerson{}).(*data.Person)
	p.logger.Info("[DEBUG] update a person with ID", person.ID)

	err := data.UpdatePerson(person)
	if err == data.ErrPersonNotFound {
		p.logger.Error("[ERROR] update a person with ID", person.ID)

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
