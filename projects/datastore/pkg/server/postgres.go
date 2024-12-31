package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PostgresApi struct {
}

func InitPostgresApi(r *mux.Router) {
	postgresApi := &PostgresApi{}
	r.HandleFunc("/postgres/persons", postgresApi.getAllPersons).Methods("GET")
	r.HandleFunc("/postgres/person", postgresApi.getPersonsBy).Methods("GET")
	r.HandleFunc("/postgres/person", postgresApi.addPerson).Methods("POST")
	r.HandleFunc("/postgres/person/{personId}", postgresApi.updatePerson).Methods("PUT")
	r.HandleFunc("/postgres/person/{personId}", postgresApi.deletePerson).Methods("DELETE")
}

func (*PostgresApi) getAllPersons(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*PostgresApi) getPersonsBy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*PostgresApi) addPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*PostgresApi) updatePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*PostgresApi) deletePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
