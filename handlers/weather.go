package handlers

import (
	"cliOpn/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// FetchWeatherData busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherData(lat, lon float64) (*models.WeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	apiUrl := os.Getenv("OPENWEATHER_API_URL")

	requestURL := fmt.Sprintf("%s?lat=%s&lon=%s&appid=%s&units=metric", apiUrl, lat, lon, apiKey)

	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch weather data: status code %d", resp.StatusCode)
	}

	var weatherData models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, fmt.Errorf("failed to decode weather data: %w", err)
	}

	return &weatherData, nil
}

// GetWeatherData é um manipulador de rota que retorna dados de previsão do tempo
func GetWeatherData(w http.ResponseWriter, r *http.Request) {

	latStr := r.URL.Query().Get("lat")
	if latStr == "" {
		http.Error(w, "O parâmetro 'lat' é obrigatório", http.StatusBadRequest)
		return
	}
	lonStr := r.URL.Query().Get("lon")
	if lonStr == "" {
		http.Error(w, "O parâmetro 'lon' é obrigatório", http.StatusBadRequest)
		return
	}
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Parâmetro 'lat' inválido", http.StatusBadRequest)
		return
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Parâmetro 'lon' inválido", http.StatusBadRequest)
		return
	}

	weatherData, err := FetchWeatherData(lat, lon)
	if err != nil {
		log.Printf("Erro ao buscar dados do tempo: %v", err)
		http.Error(w, "Erro ao buscar dados do tempo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		log.Printf("Erro ao codificar dados do tempo: %v", err)
		http.Error(w, "Erro ao codificar dados do tempo", http.StatusInternalServerError)
		return
	}
}
