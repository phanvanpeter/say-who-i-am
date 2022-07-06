package handlers

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/phanvanpeter/say-who-i-am/people/data"
	"net/http"
)

// ValidatePerson validates the person in the request and calls next if ok
func (p *People) ValidatePerson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		person := &data.Person{}
		err := data.FromJSON(person, r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid data of a person: %s", err), http.StatusBadRequest)
			return
		}

		//validate the person
		validate := validator.New()
		err = validate.Struct(person)
		if err != nil {
			// TODO return the errors to the client
			_ = err.(validator.ValidationErrors)
			http.Error(w, "Error validating the person", http.StatusUnprocessableEntity)
			return
		}

		ctx := context.WithValue(r.Context(), KeyPerson{}, person)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
