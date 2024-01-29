package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting HTTP API server")
	r := mux.NewRouter().StrictSlash(true)
	// Define r router.
	// When StrictSlash is set to true, the router will redirect requests that have a trailing slash to the same route without the trailing slash, and vice versa.
	// HandleFunc has two arguments, first URL and second function to handle it

	// GET Endpoint
	r.HandleFunc("/airports", getAirports).Methods("GET")
	r.HandleFunc("/airlines", getAirlines).Methods("GET")
	r.HandleFunc("/flights", getFlights).Methods("GET")
	r.HandleFunc("/flights/{id}", getFlightById).Methods("GET")

	// POST Endpoint
	r.HandleFunc("/flights", addFlight).Methods("POST")

	// PUT Endpoint
	r.HandleFunc("/flights/{id}", updateFlight).Methods("PUT")

	// DELETE Endpoint
	r.HandleFunc("/flights/{id}", deleteFlight).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
