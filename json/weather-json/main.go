package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var url string = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true&hourly=temperature_2m,relativehumidity_2m,windspeed_10m"

type (
	wResult struct {
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		CurrentWeather `json:"current_weather"`
	}

	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WindSpeed   float64 `json:"windspeed"`
	}
)

func main() {

	resp, err := http.Get(url)
	if err != nil {

		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Decode the JSON response into our struct type.
	var wr wResult
	err = json.NewDecoder(resp.Body).Decode(&wr)

	fmt.Println("Lattitude: ", wr.Latitude)
	fmt.Println("Longitude: ", wr.Longitude)
	fmt.Println("Temperature: ", wr.Temperature)
	fmt.Println("WindSpeed: ", wr.WindSpeed)
}
