package main

import "time"

var airline1 = Airline{
	ID: "AA", Name: "American Airlines",
}
var airline2 = Airline{
	ID: "UA", Name: "United Airlines",
}
var airlines = []Airline{
	airline1, airline2,
}

var airport1 = Airport{
	ID: "JFK", Name: "John F. Kennedy International Airport", Location: "New York",
}
var airport2 = Airport{
	ID: "LAX", Name: "Los Angeles International Airport", Location: "Los Angeles",
}
var airport3 = Airport{
	ID: "SFO", Name: "San Francisco International Airport", Location: "San Francisco",
}
var airports = []Airport{
	airport1, airport2, airport3,
}

var flights = []Flight{
	{
		ID:             "AA123",
		Airline:        airline1,
		DepartureCity:  airport1,
		ArrivalCity:    airport2,
		DepartureTime:  time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		ArrivalTime:    time.Now().Add(26 * time.Hour).Format(time.RFC3339),
		Price:          250.50,
		AvailableSeats: 120,
	},
	{
		ID:             "UA456",
		Airline:        airline2,
		DepartureCity:  airport3,
		ArrivalCity:    airport1,
		DepartureTime:  time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		ArrivalTime:    time.Now().Add(50 * time.Hour).Format(time.RFC3339),
		Price:          320.75,
		AvailableSeats: 90,
	},
	{
		ID:             "AA789",
		Airline:        airline1,
		DepartureCity:  airport2,
		ArrivalCity:    airport3,
		DepartureTime:  time.Now().Add(72 * time.Hour).Format(time.RFC3339),
		ArrivalTime:    time.Now().Add(74 * time.Hour).Format(time.RFC3339),
		Price:          200.25,
		AvailableSeats: 150,
	},
}
