package output

import (
	"cliOpn/models"
	"fmt"
)

// formatWeatherData formata os dados de previsão do tempo para exibição
// Recebe os dados como interface{} e retorna uma string formatada
type WeatherDataFormatted struct {
	Temp          float64 `json:"temp"`
	Humidity      int     `json:"humidity"`
	Clouds        int     `json:"clouds"`
	Precipitation float64 `json:"precipitation"`
	Rain          float64 `json:"rain"`
}

func FormatWeatherData(resp models.WeatherResponse) WeatherDataFormatted {
	fmt.Println("Formatting weather data...")

	return WeatherDataFormatted{
		Temp:          resp.Current.Temp,
		Humidity:      resp.Current.Humidity,
		Clouds:        resp.Current.Clouds,
		Precipitation: getPrecipitation(resp.Minutely),
		Rain:          getRain(resp.Daily),
	}
}

func getPrecipitation(minutely []models.Minutely) float64 {
	if len(minutely) > 0 {
		return minutely[0].Precipitation
	}
	return 0
}

func getRain(daily []models.Daily) float64 {
	if len(daily) > 0 {
		return daily[0].Rain
	}
	return 0
}
