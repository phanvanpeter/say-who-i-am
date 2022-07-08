package reviewshandlers

import (
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strconv"
)

type Reviews struct {
	logger hclog.Logger
}

// KeyReview is a key for passing the review through the request context
type KeyReview struct{}

// NewReviews initiliazes handler for reviews
func NewReviews(logger hclog.Logger) *Reviews {
	return &Reviews{logger: logger}
}

// GetReviewID gets the ID of the reviews to be found in the URI: /reviews/{id:[0-9]+}
func GetReviewID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}
