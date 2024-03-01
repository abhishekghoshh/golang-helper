package controllers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

const JSON_FILE_LOCATION = "controllers/persons.json"

type Person struct {
	Id        int    `json:"id" xml:"id"`
	FirstName string `json:"firstName" xml:"firstName"`
	LastName  string `json:"lastName" xml:"lastName"`
	Age       int    `json:"age" xml:"age"`
	Gender    string `json:"gender" xml:"gender"`
}

var persons []Person = []Person{
	{Id: 1, FirstName: "Abhishek", LastName: "Ghosh", Age: 25, Gender: "Male"},
	{Id: 2, FirstName: "Nasim", LastName: "Molla", Age: 27, Gender: "Male"},
	{Id: 3, FirstName: "Bishal", LastName: "Mukherjee", Age: 27, Gender: "Male"},
	{Id: 4, FirstName: "Abhishek", LastName: "Pal", Age: 26, Gender: "Male"},
}

func init() {
	if personsRetrieved, err := fetchPersonFromDisk(); err == nil {
		persons = personsRetrieved
	}
}
func cleanup() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		savePersonsToDisk()
		os.Exit(0)
	}()
}

func fetchPersonFromDisk() ([]Person, error) {
	var personsRetrieved []Person
	if bs, err := read(JSON_FILE_LOCATION); err != nil {
		return nil, err
	} else if len(bs) == 0 {
		return nil, errors.New("file is empty")
	} else {
		if err := json.Unmarshal(bs, &personsRetrieved); err != nil {
			return nil, err
		}
	}
	return personsRetrieved, nil
}

func savePersonsToDisk() {
	data, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	if err = json.Indent(&out, data, "", "  "); err != nil {
		panic(err)
	}
	if err = save(JSON_FILE_LOCATION, out.Bytes()); err != nil {
		panic(err)
	}
}

func save(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0666); err != nil {
		return err
	}
	return nil
}

func read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func GetPersons() []Person {
	return persons
}

func GetPerson(id int) Person {
	for _, p := range persons {
		if p.Id == id {
			return p
		}
	}
	return Person{}
}

func AddPerson(person *Person) *Person {
	n := len(persons)
	person.Id = persons[n-1].Id + 1
	persons = append(persons, *person)
	return person
}

func AddToResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "application/xml":
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(data)
	case "application/json":
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	default:
		t := reflect.TypeOf(data)
		if t.Kind() == reflect.Struct || t.Kind() == reflect.Array || t.Kind() == reflect.Map || t.Kind() == reflect.Slice {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)
		} else {
			fmt.Fprintf(w, data.(string))
		}
	}
}
