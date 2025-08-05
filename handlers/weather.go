package handlers

import (
	"cliOpn/config"
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

// FetchWeatherDataWithCoordinates busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherDataWithCoordinates(lat, lon float64) (*models.WeatherResponse, error) {
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

// FetchWeatherDataWithJson busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros de um JSOn da aplicação e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherDataWithJson() (*models.WeatherResponse, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Erro ao obter configuração: %v", err)
		return nil, err
	}
	return FetchWeatherDataWithCoordinates(cfg.Lat, cfg.Lon)
}

// GetWeatherData é um manipulador de rota que retorna dados de previsão do tempo
func GetWeatherData(w http.ResponseWriter, r *http.Request) {
	var weatherData *models.WeatherResponse
	var err error

	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	if latStr != "" && lonStr != "" {
		lat, err := strconv.ParseFloat(latStr, 64)
		lon, err := strconv.ParseFloat(lonStr, 64)
		if err != nil {
			http.Error(w, "Parâmetros 'lat' ou 'lon' inválidos", http.StatusBadRequest)
			return
		}
		weatherData, err = FetchWeatherDataWithCoordinates(lat, lon)
	} else {
		weatherData, err = FetchWeatherDataWithJson()
		if err != nil {
			log.Printf("Erro ao buscar dados do tempo: %v", err)
			http.Error(w, "Erro ao buscar dados do tempo", http.StatusInternalServerError)
			return
		}
		if weatherData == nil {
			http.Error(w, "Dados do tempo não encontrados", http.StatusNotFound)
			return
		}
	}
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
