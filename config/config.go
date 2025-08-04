package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	ExcludedFields []string `json:"exclude"`
}

const configPath = "config.json"

var AppConfig Config

func LoadConfig() (*Config, error) {
	// Configuração padrão
	conf := &Config{
		Lat:            40.7128,  // Exemplo: Nova York
		Lon:            -74.0060, // Exemplo: Nova York
		ExcludedFields: []string{"minutely", "hourly", "daily", "alerts"},
	}

	data, err := os.ReadFile(configPath)
	// Se o arquivo não existe, podemos ignorar o erro e usar o padrão.
	// Para outros erros, retornamos o erro.
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Se o arquivo existe, decodifica o JSON para a struct
	if err == nil {
		err = json.Unmarshal(data, conf)
		if err != nil {
			return nil, err
		}
	}

	return conf, nil
}
func (c *Config) SaveConfig() error {
	// MarshalIndent formata o JSON para ser legível
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	// Escreve os dados no arquivo, sobrescrevendo o conteúdo existente.
	// 0644 são as permissões do arquivo.
	return os.WriteFile(configPath, data, 0644)
}

func UpdateCoordinates(lat, lon float64) error {
	conf, err := LoadConfig()
	if err != nil {
		return err
	}

	conf.Lat = lat
	conf.Lon = lon

	return conf.SaveConfig()
}

func UpdatexcludedFields(excluded []string) error {
	conf, err := LoadConfig()
	if err != nil {
		return err
	}

	conf.ExcludedFields = excluded

	return conf.SaveConfig()
}
