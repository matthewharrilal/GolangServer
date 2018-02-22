package main

// So let us tackle this from the beginning and what is happening
import (
	// These are all the libraries that we are importing
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

// We are creating a Pokemon struct and these are the attributes that the class will have
type Pokemon struct {
	Name string "json:'name,omitempty'"
	PowerLevel int "json: 'powerLevel,omitempty'"
	Type string "json:'type,omitempty'"
}

// We are initializing an empty array ok The Object Pokemon
var pokemon []Pokemon

// Now when we compare this to when we created servers in Flask this is the function that will fetch the pokemon and
// what tell us that this function is going to be actually making the get request is in the main function below
func getPokemon(w http.ResponseWriter, req *http.Request) {
	// The reason that we are posting the users  inside the get function is because we do not have a database to store
	// the posted users so we post them and fetch the same results without returning to quickly see if our routes are
	// working
	json.NewEncoder(w).Encode(pokemon)
}

// We are not going to use this function because we do not have a database that we can post the resources to as well as
// fetch them
func postPokemon(w http.ResponseWriter, req *http.Request) {

}

func main() {
	// We are creating a new router which is essentially urlsession.shared in swift
	router := mux.NewRouter()

	// We are appending to the global array that we made of type Pokemon and we are giving that pokemon a the given
	// attributes as well as the type and the reason that we are doing that is because we want to be able to use dummy
	// data since we do not have a database to work with
	pokemon = append(pokemon, Pokemon{("Bubbles"), 12, "Water type"})

	// In this function right below we are configuring our routes tob handle the showPokemon route which will return all
	// the available pokemon and we configure with the GET Method which we have declared
	router.HandleFunc("/showPokemon", getPokemon).Methods("GET")
	// And this starts listening to the port to establish the connection
	log.Fatal(http.ListenAndServe(":12345", router))
}