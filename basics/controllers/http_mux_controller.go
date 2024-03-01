package controllers

import (
	"log"
	"net/http"
)

func HttpMuxController() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, "Hi this is http mux controller")
		},
	)

	mux.HandleFunc(
		"/persons",
		func(w http.ResponseWriter, r *http.Request) {
			AddToResponse(w, r, GetPersons())
		},
	)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
