package data

import "fmt"

var ErrPersonNotFound = fmt.Errorf("person not found")

// Person is a struct for person
type Person struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"firstName" validate:"required"`
	LastName    string  `json:"lastName" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Age         int     `json:"age" validate:"required,gte=15,lte=115"`
	Description string  `json:"description"`
	Stars       float32 `json:"stars" validate:"required,gte=1.0,lte=5.0"`
}

// People is an "alias" for the list of persons.
type People []*Person

// peopleList is a dummy storage for the people (list of persons).
var peopleList = []*Person{
	&Person{
		ID:          1,
		FirstName:   "Mark",
		LastName:    "Lewis",
		Age:         45,
		Description: "Mark Lewis is not the child of the book writer C.S.Lewis",
		Stars:       4.5,
	},
	&Person{
		ID:          2,
		FirstName:   "Donald",
		LastName:    "Trump",
		Age:         75,
		Description: "The former president of the US",
		Stars:       3.7,
	},
}
