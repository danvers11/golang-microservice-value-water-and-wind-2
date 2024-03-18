package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

type Response struct {
	Wind   float64 `json:"wind"`
	Water  float64 `json:"water"`
	Status string  `json:"status"`
}

func main() {
	http.HandleFunc("/update", updateHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	wind := rand.Float64() * 20  // Generate random wind value
	water := rand.Float64() * 10 // Generate random water value

	status := getStatus(wind, water)

	response := Response{
		Wind:   wind,
		Water:  water,
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getStatus(wind, water float64) string {
	var status string

	if water < 5 || wind < 6 {
		status = "Aman"
	} else if water >= 6 && water <= 8 || wind >= 7 && wind <= 15 {
		status = "Siaga"
	} else {
		status = "Bahaya"
	}

	return status
}
