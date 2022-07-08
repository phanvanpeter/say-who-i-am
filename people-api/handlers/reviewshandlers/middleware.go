package reviewshandlers

import (
	"context"
	"fmt"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

// ValidateReview validates the review in the body of the request
func (rv *Reviews) ValidateReview(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rv.logger.Info("Validate Review middleware")

		review := data.Review{}
		err := data.FromJSON(&review, r.Body)
		if err != nil {
			rv.logger.Error("Error deserializing request body: review")
			http.Error(w, fmt.Sprintf("invalid request body, expected review: %s", err), http.StatusBadRequest)
			return
		}

		// TODO validate

		ctx := context.WithValue(r.Context(), KeyReview{}, review)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
