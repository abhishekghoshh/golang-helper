package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GorillaMuxController() {
	cleanup()
	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, "Hi this is Gorilla mux controller")
		},
	)

	router.HandleFunc(
		"/persons",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, GetPersons())
		},
	).Methods("GET")

	router.HandleFunc(
		"/persons/{personId:[0-9]+}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			personId, _ := strconv.Atoi(vars["personId"])
			onePerson := GetPerson(personId)
			AddToResponse(w, r, onePerson)
		},
	).Methods("GET")

	router.HandleFunc(
		"/persons",
		func(w http.ResponseWriter, r *http.Request) {
			var onePerson Person
			err := json.NewDecoder(r.Body).Decode(&onePerson)
			if err != nil {
				panic(err)
			}
			AddPerson(&onePerson)
			AddToResponse(w, r, onePerson)
		},
	).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
