package service

import (
	"basic-controller/person"
	"encoding/json"
)

var persons []person.Person

func GetPersons() []person.Person {
	if len(persons) == 0 {
		json.Unmarshal(Retrieve("persons.json"), &persons)
	}
	return persons
}

func GetPerson(id int) person.Person {
	var person person.Person
	persons := GetPersons()
	for _, p := range persons {
		if p.Id == id {
			return p
		}
	}
	return person
}

func AddPerson(person *person.Person) *person.Person {
	GetPersons()
	length := len(persons)
	person.Id = length + 1
	persons = append(persons, *person)
	return person
}
