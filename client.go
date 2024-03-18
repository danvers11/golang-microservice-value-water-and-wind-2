package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Wind   float64 `json:"wind"`
	Water  float64 `json:"water"`
	Status string  `json:"status"`
}

func main() {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			update()
		}
	}
}

func update() {
	resp, err := http.Get("http://localhost:8081/update")
	if err != nil {
		fmt.Println("Failed to fetch data:", err)
		return
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("Failed to decode response:", err)
		return
	}

	fmt.Printf("Wind: %.2f m/s, Water: %.2f m, Status: %s\n", response.Wind, response.Water, response.Status)
}
