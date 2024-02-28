package controller

import (
	"basic-controller/person"
	"basic-controller/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc("/persons", getPersons).Methods(http.MethodGet)
	router.HandleFunc("/persons/{personId:[0-9]+}", getPerson).Methods(http.MethodGet)
	router.HandleFunc("/persons", addPerson).Methods(http.MethodPost)
}
func getPersons(w http.ResponseWriter, r *http.Request) {
	persons := service.GetPersons()
	service.AddToResponse(w, r, persons)
}
func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personId, _ := strconv.Atoi(vars["personId"])
	onePerson := service.GetPerson(personId)
	service.AddToResponse(w, r, onePerson)
}
func addPerson(w http.ResponseWriter, r *http.Request) {
	var onePerson person.Person
	err := json.NewDecoder(r.Body).Decode(&onePerson)
	if err != nil {
		panic(err)
	}
	service.AddPerson(&onePerson)
	service.AddToResponse(w, r, onePerson)
}
