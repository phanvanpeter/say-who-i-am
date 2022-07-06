package data

type Person struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Age         int     `json:"age"`
	Description string  `json:"description"`
	Stars       float32 `json:"stars"`
}

type People []*Person

func GetPeople() People {
	return peopleList
}

func AddPerson(p *Person) {
	peopleList = append(peopleList, p)
}

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
