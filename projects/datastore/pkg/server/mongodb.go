package server

import (
	"net/http"

	"github.com/abhishekghoshh/datastore/pkg/mongodb"
	"github.com/gorilla/mux"
)

type MongoDBApi struct {
	mongoDb *mongodb.MongoDb
}

func NewMongoDbApi(mongoDb *mongodb.MongoDb) *MongoDBApi {
	return &MongoDBApi{mongoDb}
}

func (api *MongoDBApi) Init(r *mux.Router) {
	r.HandleFunc("/mongo/persons", api.getAllPersons).Methods("GET")
	r.HandleFunc("/mongo/person", api.getPersonsBy).Methods("GET")
	r.HandleFunc("/mongo/person", api.addPerson).Methods("POST")
	r.HandleFunc("/mongo/person/{personId}", api.updatePerson).Methods("PUT")
	r.HandleFunc("/mongo/person/{personId}", api.deletePerson).Methods("DELETE")
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
