package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)

type Person struct {
	ID string "json:'id,omitempty'"
	Firstname string "json:'firstname,omitempty'"
	Lastname string "json:'lastname,omitempty'"
	// Since the value for the key address is a dictionary we have to reference the other struct like this
	Address *Address "json:'address,omitempty'"
}

type Address struct {
	City string "json:'city,omitempty'"
	State string "json:'state,omitempty'"
}

var people []Person

func getPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	// We are taking the variable that we made global that encodes the array of people
	json.NewEncoder(w).Encode(people)
}

func createPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func deletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	// Right now we are declaring the router
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Matthew", Lastname: "Harrilal", Address: &Address{City: "New York", State:"California"}})
	// Configuring our routes
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("people/{id}", getPersonEndpoint).Methods("GET")
	router.HandleFunc("people/{id}", createPersonEndpoint).Methods("POST")
	router.HandleFunc("people/{id}", deletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
