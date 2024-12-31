package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MongoDBApi struct {
}

func InitMongoDBApi(r *mux.Router) {
	mongoApi := &MongoDBApi{}
	r.HandleFunc("/mongo/persons", mongoApi.getAllPersons).Methods("GET")
	r.HandleFunc("/mongo/person", mongoApi.getPersonsBy).Methods("GET")
	r.HandleFunc("/mongo/person", mongoApi.addPerson).Methods("POST")
	r.HandleFunc("/mongo/person/{personId}", mongoApi.updatePerson).Methods("PUT")
	r.HandleFunc("/mongo/person/{personId}", mongoApi.deletePerson).Methods("DELETE")
}

func (*MongoDBApi) getAllPersons(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*MongoDBApi) getPersonsBy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*MongoDBApi) addPerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*MongoDBApi) updatePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (*MongoDBApi) deletePerson(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
