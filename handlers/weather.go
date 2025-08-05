package handlers

import (
	"cliOpn/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// FetchWeatherData busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherData(lat, lon float64) (*models.WeatherResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	apiUrl := os.Getenv("OPENWEATHER_API_URL")

	ApiUrl := fmt.Sprintf("%s?lat=%f&lon=%f&appid=%s", apiUrl, lat, lon, apiKey)

	fmt.Println(ApiUrl)

	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, ApiUrl, payload)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	fmt.Println(string(body))
	// Arrumar depois para ajustar a leitura do JSON de resposta do CLI
	return &models.WeatherResponse{}, nil
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
