package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Weather struct {
	Description string `json:"description"`
	Temperature int `json:"temperature"` 
}
func getWeather(city string) (*Weather, error) {
	url := fmt.Sprintf('http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=API_KEY', city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var weather weather
	err = json.NewDecoder(resp.body).Decode(&weather)
	if err != nil {
		return nil, err
	}
	return weather, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: weatherapp <city>")
		os.Exit(1)
	}
	city := os.Args[1]
	weather, err := getWeather(city)
	if err != nil {
		fmt.Println("Failed to get weather")
		os.Exit(1)
	}
	fmt.Println("Weather for ", city+":")
	fmt.Println("description: ", weather.Description)
	fmt.Println("Temp: ", weather.Temperature, "°C")

}