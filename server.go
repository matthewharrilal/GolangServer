package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
)

// First things first let us model the data that we are to go and try to post to the database
type Pokemon struct {
	Name string "json:'name,omitempty'"
	PowerLevel int "json: 'powerLevel,omitempty'"
	Type string "json:'type,omitempty'"
}

var pokemon []Pokemon
func getPokemon(w http.ResponseWriter, req *http.Request) {
	// When we the user makes a get request to the server we make a post request of the pokemon and therefore we can
	// fetch the pokemon
	json.NewEncoder(w).Encode(pokemon)
}

func postPokemon(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	pokemon = append(pokemon, Pokemon{("Bubbles"), 12, "Water type"})
	router.HandleFunc("/showPokemon", getPokemon).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}