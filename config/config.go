package config

import (
	"encoding/json"
	"os"
	"sync"
)

type Config struct {
	Lat            float64  `json:"lat"`
	Lon            float64  `json:"lon"`
	ExcludedFields []string `json:"exclude"`
}

const configPath = "Json/config.json"

var (
	instance *Config
	once     sync.Once
	loadErr  error
)

func GetConfig() (*Config, error) {
	once.Do(func() {
		// Configuração padrão
		conf := &Config{
			Lat:            40.7128,  // Exemplo: Nova York
			Lon:            -74.0060, // Exemplo: Nova York
			ExcludedFields: []string{"minutely", "hourly", "daily", "alerts"},
		}

		data, err := os.ReadFile(configPath)
		if err != nil {

			if os.IsNotExist(err) {
				instance = conf
				return
			}

			loadErr = err
			return
		}

		if err := json.Unmarshal(data, conf); err != nil {
			loadErr = err
			return
		}
		instance = conf
	})

	return instance, loadErr
}

func (c *Config) SaveConfig() error {

	dirPath := "Json"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}

	configPath := dirPath + "/config.json"

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

func UpdateCoordinates(lat, lon float64) error {
	conf, err := GetConfig()
	if err != nil {
		return err
	}

	conf.Lat = lat
	conf.Lon = lon

	return conf.SaveConfig()
}

func UpdatexcludedFields(excluded []string) error {
	conf, err := GetConfig()
	if err != nil {
		return err
	}

	conf.ExcludedFields = excluded

	return conf.SaveConfig()
}
