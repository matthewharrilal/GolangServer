package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

// First thing first we are essentially declaring a schema for how the data we are going to send and fetch is going to
// look

type Human struct {
	// We have created a human struct which contains the following values
	Firstname string "json:'firstname,omitempty'"
	Lastname string "json:'lastname,omitempty'"
	Age int "json:'age,omitempty'"
}


// Now since we do not have a database currently we are going to create a global array which is going to contain dummy
// data about the human we are going to use to our manipulation
var sampleHumans  []Human

func showSampleHumans(response http.ResponseWriter, req *http.Request) {
	// So what this function is essentially going to do for us is that it is going to show us all the humans that we
	// have in our collection

	// What this line of code does for us is that it allows us to encode resources in the response and this will be sent
	// over the network when this function is executed and the reason that we do this is because we do not have a data
	// base to send these over to so we have to work with what we got
	json.NewEncoder(response).Encode(sampleHumans)
}

func main() {
	// So here we are declaring a new router which is similar to how in swift we essentially declaring a url session
	var router = mux.NewRouter()
	sampleHumans = append(sampleHumans, Human{"Matthew", "Harrilal", 12})
	router.HandleFunc("/showHumans",showSampleHumans).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}