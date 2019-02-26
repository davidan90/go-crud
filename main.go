package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Address   Address
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "David", Address: Address{City: "Madrid", State: "Spain"}})
	people = append(people, Person{ID: "2", FirstName: "Maria", Address: Address{City: "Madrid", State: "Spain"}})

	// endpoints
	router.HandleFunc("/people", GetAllPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetAllPeopleEndpoint(resp http.ResponseWriter, req *http.Request) {
	json.NewEncoder(resp).Encode(people)
}

func GetPeopleEndpoint(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}
	}
	json.NewEncoder(resp).Encode(Person{})
}

func CreatePeopleEndpoint(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)

	json.NewEncoder(resp).Encode(people)
}

func DeletePeopleEndpoint(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range people {
		if item.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(resp).Encode(people)
}
