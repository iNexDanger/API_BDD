package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRoutesForPersons(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		persons, err := getPersons()
		if err == nil {
			respondWithSuccess(persons, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		videogame, err := getPersonById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(videogame, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createPerson(person)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		// Declare a var so we can decode json into it
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updatePersons(person)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		err = deletePerson(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}
