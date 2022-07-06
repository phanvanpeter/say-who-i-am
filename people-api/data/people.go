package data

// GetPeople returns all the people from the storage.
func GetPeople() People {
	return peopleList
}

func GetPerson(id int) (*Person, error) {
	i := findPerson(id)
	if i == -1 {
		return nil, ErrPersonNotFound
	}
	return peopleList[i], nil
}

// AddPerson adds a person to the storage.
func AddPerson(p *Person) {
	p.ID = getNextID()
	peopleList = append(peopleList, p)
}

// UpdatePerson updates a person with a given ID in the storage
func UpdatePerson(p *Person) error {
	i := findPerson(p.ID)
	if i == -1 {
		return ErrPersonNotFound
	}
	peopleList[i] = p
	return nil
}

// DeletePerson deletes a person according the ID
func DeletePerson(id int) error {
	i := findPerson(id)
	if i == -1 {
		return ErrPersonNotFound
	}

	peopleList = append(peopleList[:i], peopleList[i+1:]...)
	return nil
}

// findPerson finds a person according to the ID.
// If the person is not in the storage, the functions returns -1
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
