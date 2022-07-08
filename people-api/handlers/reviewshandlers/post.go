package reviewshandlers

import (
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

// Create handles creating a new review in the storage
func (rv *Reviews) Create(w http.ResponseWriter, r *http.Request) {
	review := r.Context().Value(KeyReview{}).(data.Review)
	rv.logger.Info("review")
	rv.logger.Info(review.Text)

	//w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
	err := data.ToJSON(review, w)
	if err != nil {
		http.Error(w, "Unable to serialize review into JSON", http.StatusInternalServerError)
	}
}
