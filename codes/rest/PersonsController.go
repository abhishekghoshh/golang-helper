package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName" xml:"firstName"`
	LastName  string `json:"lastName" xml:"lastName"`
	Age       int    `json:"age" xml:"age"`
	Gender    string `json:"gender" xml:"gender"`
}

var persons []Person

func main() {
	http.HandleFunc("/persons", getPersons)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func getPersons(w http.ResponseWriter, r *http.Request) {
	persons = []Person{
		{FirstName: "Abhishek", LastName: "Ghosh", Age: 25, Gender: "Male"},
		{FirstName: "Nasim", LastName: "Molla", Age: 26, Gender: "Male"},
		{FirstName: "Bishal", LastName: "Molla", Age: 26, Gender: "Male"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(persons)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}

}
