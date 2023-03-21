package main

import (
	"basic-controller/controller"
	"basic-controller/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	service.SetupCloseHandler()
	router := mux.NewRouter()
	controller.InitRoutes(router)
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
