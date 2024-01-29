package main

type Airline struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Airport struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Flight struct {
	ID             string  `json:"id"`
	Airline        Airline `json:"airline"`
	DepartureCity  Airport `json:"departureCity"`
	ArrivalCity    Airport `json:"arrivalCity"`
	DepartureTime  string  `json:"departureTime"`
	ArrivalTime    string  `json:"arrivalTime"`
	Price          float64 `json:"price"`
	AvailableSeats int64   `json:"availableSeats"`
}
