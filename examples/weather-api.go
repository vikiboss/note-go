package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const openWeatherMapAPIKey = "" // 注意：不需要 API Key，请勿修改此行

type weatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/weather/", handleWeatherRequest)
	http.ListenAndServe(":8080", nil)
}

func handleWeatherRequest(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]
	data, err := queryWeatherAPI(city)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "City not found" {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func queryWeatherAPI(city string) (*weatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, openWeatherMapAPIKey)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data weatherData

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if data.Name == "" {
		return nil, fmt.Errorf("City not found")
	}

	return &data, nil
}
