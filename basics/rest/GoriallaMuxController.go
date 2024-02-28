package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rest/core"
	"rest/person"
	"strconv"

	"github.com/gorilla/mux"
)

func Main() {
	router := mux.NewRouter()
	router.HandleFunc("/persons", getPersons).Methods(http.MethodGet)
	router.HandleFunc("/persons/{personId:[0-9]+}", getPerson).Methods(http.MethodGet)
	router.HandleFunc("/persons", addPerson).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
func getPersons(w http.ResponseWriter, r *http.Request) {
	persons := person.GetPersons()
	core.AddToResponse(w, r, persons)
}
func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])
	onePerson := person.GetPerson(personId)
	core.AddToResponse(w, r, onePerson)
}
func addPerson(w http.ResponseWriter, r *http.Request) {
	var onePerson person.Person
	err := json.NewDecoder(r.Body).Decode(&onePerson)
	if err != nil {
		panic(err)
	}
	person.AddPerson(&onePerson)
	core.AddToResponse(w, r, onePerson)
}
