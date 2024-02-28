package person

var persons []Person

func GetPersons() []Person {
	if len(persons) == 0 {
		persons = []Person{
			{Id: 1, FirstName: "Abhishek", LastName: "Ghosh", Age: 25, Gender: "Male"},
			{Id: 2, FirstName: "Nasim", LastName: "Molla", Age: 26, Gender: "Male"},
			{Id: 3, FirstName: "Bishal", LastName: "Molla", Age: 26, Gender: "Male"},
		}
	}
	return persons
}

func GetPerson(id int) Person {
	var person Person
	persons := GetPersons()
	for _, p := range persons {
		if p.Id == id {
			return p
		}
	}
	return person
}

func AddPerson(person *Person) *Person {
	GetPersons()
	length := len(persons)
	person.Id = length + 1
	persons = append(persons, *person)
	return person
}
