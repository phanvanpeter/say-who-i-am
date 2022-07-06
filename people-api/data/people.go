package data

type Person struct {
	ID          int
	FirstName   string
	LastName    string
	Age         int
	Description string
	Starts      float32
}

type People []*Person

func GetPeople() People {
	return peopleList
}

var peopleList = []*Person{
	&Person{
		ID:          1,
		FirstName:   "Mark",
		LastName:    "Lewis",
		Age:         45,
		Description: "Mark Lewis is not the child of the book writer C.S.Lewis",
		Starts:      4.5,
	},
	&Person{
		ID:          2,
		FirstName:   "Donald",
		LastName:    "Trump",
		Age:         75,
		Description: "The former president of the US",
		Starts:      3.7,
	},
}
