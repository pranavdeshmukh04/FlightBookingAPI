package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func getAirports(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Airports Data")
	json.NewEncoder(w).Encode(airports)
}
func getAirlines(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Airlines Data")
	json.NewEncoder(w).Encode(airlines)
}
func getFlights(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Flights data")
	json.NewEncoder(w).Encode(flights)
}

func getFlightById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Fetching Flight By Id:", id)

	for _, flight := range flights {
		if flight.ID == id {
			json.NewEncoder(w).Encode(flight)
		}
	}
}

func addFlight(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding Flight")
	reqBody, _ := io.ReadAll(r.Body)

	var flight Flight
	json.Unmarshal(reqBody, &flight)
	// This line unmarshals (or parses) the JSON-encoded data from the request
	// body into a Grocery struct.
	flights = append(flights, flight)
	// the grocery item is appended to the groceries slice.
	json.NewEncoder(w).Encode(flights)
}

func deleteFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Deleting Flight By Id:", id)

	for index, flight := range flights {
		if flight.ID == id {
			flights = append(flights[:index], flights[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(flights)
}

func updateFlight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Updating Flight By Id:", id)
	for index, flight := range flights {
		if flight.ID == id {
			flights = append(flights[:index], flights[index+1:]...)
			var updateflight Flight

			json.NewDecoder(r.Body).Decode(&updateflight)
			// decodes the JSON data from the request body into a Grocery struct
			// representing the updated grocery item.
			flights = append(flights, updateflight)

			json.NewEncoder(w).Encode(updateflight)
			return
		}
	}
}
