package handlers

import (
	"cliOpn/config"
	"cliOpn/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func loadConfig() {
	// Encontra o diretório 'home' do usuário
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Erro: não foi possível encontrar o diretório home do usuário: %v", err)
	}

	// Monta o caminho para o arquivo de configuração padrão
	configPath := filepath.Join(homeDir, ".config", "weatherapp", "config")

	// Tenta carregar o arquivo de configuração.
	// Se o arquivo não existir, godotenv.Load ignora o erro silenciosamente.
	// Isso permite que as variáveis de ambiente do sistema ainda funcionem como um fallback.
	godotenv.Load(configPath)
}

// FetchWeatherDataWithCoordinates busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherDataWithCoordinates(lat, lon float64, exclude []string) (*models.WeatherResponse, error) {
	// Carrega as configurações do arquivo padrão ou do ambiente
	loadConfig()

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	apiUrl := os.Getenv("OPENWEATHER_API_URL")

	if apiKey == "" {
		return nil, fmt.Errorf("a variável de ambiente OPENWEATHER_API_KEY não está definida. Execute o script de setup ou exporte a variável")
	}

	formatExcluded := strings.Join(exclude, ",")
	ApiUrl := fmt.Sprintf("%s?lat=%f&lon=%f&exclude=%s&units=metric&lang=pt_br&appid=%s", apiUrl, lat, lon, formatExcluded, apiKey)

	res, err := http.Get(ApiUrl)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Erro ao fechar o corpo da resposta: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API retornou um status não-OK: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	var weatherResponse models.WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &weatherResponse, nil
}

// FetchWeatherDataWithJson busca dados de previsão do tempo da API OpenWeather Recebe latitude e longitude como parâmetros de um JSOn da aplicação e retorna uma estrutura WeatherResponse ou um erro
func FetchWeatherDataWithJson() (*models.WeatherResponse, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Erro ao obter configuração: %v", err)
		return nil, err
	}
	return FetchWeatherDataWithCoordinates(cfg.Lat, cfg.Lon, cfg.ExcludedFields)
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
		cfg, err := config.GetConfig()

		weatherData, err = FetchWeatherDataWithCoordinates(lat, lon, cfg.ExcludedFields)
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
