package peoplehandlers

import (
	"context"
	"fmt"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

// ValidatePerson validates the person in the request and calls next if ok
func (p *People) ValidatePerson(next http.Handler) http.Handler {
	p.logger.Info("[DEBUG] update a person")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		person := &data.Person{}
		err := data.FromJSON(person, r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid data of a person: %s", err), http.StatusBadRequest)
			return
		}

		//validate
		errs := p.validate.Validate(person)
		if errs != nil {
			p.logger.Error("error validating a person")

			w.WriteHeader(http.StatusUnprocessableEntity)
			d := ValidationError{
				Messages: errs.Errors(),
			}

			_ = data.ToJSON(d, w)
			return
		}

		ctx := context.WithValue(r.Context(), KeyPerson{}, person)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
