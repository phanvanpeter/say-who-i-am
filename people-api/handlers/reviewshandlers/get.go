package reviewshandlers

import (
	"fmt"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

func (rv *Reviews) GetAll(w http.ResponseWriter, r *http.Request) {
	rv.logger.Info("Get all reviews")

	reviews := data.GetReviews()
	err := data.ToJSON(reviews, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to fetch the reviews: %s", err), http.StatusInternalServerError)
	}
}

func (rv *Reviews) GetByID(w http.ResponseWriter, r *http.Request) {
	rv.logger.Info("Get review by ID")

	id, err := GetReviewID(r)
	if err != nil {
		rv.logger.Error("Invalid ID of review in URI", err)
		http.Error(w, fmt.Sprintf("Invalid ID of review in URI: %s", err), http.StatusBadRequest)
		return
	}

	review, err := data.GetReview(id)
	if err != nil {
		rv.logger.Error("Review not found", err)
		http.Error(w, string(fmt.Sprintf("Review with ID %d not found: %s\n", id, err)), http.StatusNotFound)
		return
	}

	err = data.ToJSON(review, w)
	if err != nil {
		rv.logger.Error("error serializing the review into JSON: %s", err)
		http.Error(w, fmt.Sprintf("error serializing the review into JSON: %s", err), http.StatusInternalServerError)
		return
	}
}
