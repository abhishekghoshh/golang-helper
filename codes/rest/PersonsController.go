package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"rest/person"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/persons", getPersons)
	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}

func getPersons(w http.ResponseWriter, r *http.Request) {
	persons := person.GetStaticPersons()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(persons)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(persons)
	}
}
