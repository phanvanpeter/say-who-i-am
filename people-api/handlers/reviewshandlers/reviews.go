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

func NewReviews(logger hclog.Logger) *Reviews {
	return &Reviews{logger: logger}
}

func GetReviewID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}
