package handlers

import (
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

// Create creates a new person from the post request and add it in the storage
func (p *People) Create(w http.ResponseWriter, r *http.Request) {
	p.logger.Info("[DEBUG] create a new person")

	person := r.Context().Value(KeyPerson{}).(*data.Person)

	data.AddPerson(person)

}
