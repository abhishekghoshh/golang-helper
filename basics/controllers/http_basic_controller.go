package controllers

import (
	"log"
	"net/http"
)

func HttpBasicController() {

	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, "Hi this is http basic controller")
		},
	)

	http.HandleFunc(
		"/persons",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, GetPersons())
		},
	)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
