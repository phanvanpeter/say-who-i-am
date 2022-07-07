package peoplehandlers

import (
	"fmt"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

// Delete deletes a person from the storage according to a given ID in the URL
func (p *People) Delete(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] delete person")

	id, err := GetPersonID(r)
	if err != nil {
		p.logger.Error("[ERROR] invalid ID of person")
		http.Error(w, "Invalid ID number", http.StatusBadRequest)
		return
	}

	err = data.DeletePerson(id)
	if err == data.ErrPersonNotFound {
		p.logger.Error("[ERROR] not existing person with ID", id)

		http.Error(w, fmt.Sprintf("Person with ID %d does not exist", id), http.StatusNotFound)
	}
}
