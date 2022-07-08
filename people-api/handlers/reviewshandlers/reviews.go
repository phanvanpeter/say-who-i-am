package reviewshandlers

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/phanvanpeter/say-who-i-am/people-api/data"
	"net/http"
)

type Reviews struct {
	logger hclog.Logger
}

func NewReviews(logger hclog.Logger) *Reviews {
	return &Reviews{logger: logger}
}

func (rv *Reviews) GetAll(w http.ResponseWriter, r *http.Request) {
	rv.logger.Info("Get all reviews")

	reviews := data.GetReviews()
	err := data.ToJSON(reviews, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("unable to fetch the reviews: %s", err), http.StatusInternalServerError)
	}
}
