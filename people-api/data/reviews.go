package data

import (
	"fmt"
	"time"
)

var ErrReviewNotFound = fmt.Errorf("review not found")

type Review struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	From      *Person   `json:"from"`
	To        *Person   `json:"to"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Reviews []*Review

func GetReviews() Reviews {
	return reviewList
}

func GetReview(id int) (*Review, error) {
	i := findReview(id)
	if i == -1 {
		return nil, ErrReviewNotFound
	}
	return reviewList[i], nil
}

func findReview(id int) int {
	for i, review := range reviewList {
		if review.ID == id {
			return i
		}
	}
	return -1
}

var reviewList = Reviews{
	&Review{
		ID:        1,
		Title:     "What a boring person",
		From:      person(1),
		To:        person(2),
		Text:      "He is just boring, nothing to talk about with him",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	&Review{
		ID:        2,
		Title:     "What a wonderful work college",
		From:      person(2),
		To:        person(1),
		Text:      "He is just great",
		CreatedAt: time.Now().AddDate(-1, -1, -1),
		UpdatedAt: time.Now().AddDate(0, 0, -1),
	},
}

func person(id int) *Person {
	p, _ := GetPerson(id)
	return p
}