package handlers

import (
	"cliOpn/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// GetWeatherData é um manipulador de rota que retorna dados de previsão do tempo
func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	if lat == "" || lon == "" {
		http.Error(w, "latitude and longitude parameter is required", http.StatusBadRequest)
		fmt.Println("blank LAT or LON information")
		return
	}
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	apiUrl := os.Getenv("OPENWEATHER_API_URL")

	requestURL := apiUrl + "?q=" + lat + "," + lon + "&appid=" + apiKey + "&units=metric"
	resp, err := http.Get(requestURL)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		fmt.Println("Error fetching weather data:", err)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(resp.Body)

		if resp.StatusCode != http.StatusOK {
			http.Error(w, "Failed to fetch weather data", resp.StatusCode)
			fmt.Println("Error fetching weather data: status code", resp.StatusCode)
		}
		var weatherData models.WeatherResponse
		if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
			http.Error(w, "Failed to decode weather data", http.StatusInternalServerError)
			fmt.Println("Error JSON decoding weather data:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(weatherData); err != nil {
			http.Error(w, "Failed to encode weather data", http.StatusInternalServerError)
			fmt.Println("Error encoding weather data to JSON:", err)

			// Log the successful fetch
			// log.Printf("Weather data fetched successfully for lat: %s, lon: %s
			w.WriteHeader(http.StatusOK)
			fmt.Println("Weather data fetched successfully for lat:", lat, "lon:", lon)
			// Return the weather data as JSON response
			w.Write([]byte(fmt.Sprintf(`{"lat": "%s", "lon": "%s", "weather": %s}`, lat, lon, weatherData)))
			fmt.Println("Weather data:", weatherData)
			return

		}
	}
}
