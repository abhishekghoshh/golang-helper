package main

import (
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}

func main() {
	http.HandleFunc("/greet", greet)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi this is Abhishek")
}
