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

// GetPeople returns all the people from the storage.
func GetPeople() People {
	return peopleList
}

// AddPerson adds a person to the storage.
func AddPerson(p *Person) {
	p.ID = getNextID()
	peopleList = append(peopleList, p)
}

func UpdatePerson(p *Person) error {
	i := findPerson(p.ID)
	if i == -1 {
		return ErrPersonNotFound
	}
	peopleList[i] = p
	return nil
}

func findPerson(id int) int {
	for i, p := range peopleList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func getNextID() int {
	lastID := peopleList[len(peopleList)-1].ID
	return lastID + 1
}

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
