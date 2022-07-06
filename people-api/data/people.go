package data

// Person is a struct for person
type Person struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Age         int     `json:"age"`
	Description string  `json:"description"`
	Stars       float32 `json:"stars"`
}

// People is an "alias" for the list of persons.
type People []*Person

// GetPeople returns all the people from the storage.
func GetPeople() People {
	return peopleList
}

// AddPerson adds a person to the storage.
func AddPerson(p *Person) {
	peopleList = append(peopleList, p)
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
