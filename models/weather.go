package models


// Structs para mapear a resposta da API de previsão do tempo

type WeatherResponse struct {
	Lat            float64    `json:"lat"`
	Lon            float64    `json:"lon"`
	Timezone       string     `json:"timezone"`
	TimezoneOffset int        `json:"timezone_offset"`
	Current        Current    `json:"current"`
	Minutely       []Minutely `json:"minutely"`
	Hourly         []Hourly   `json:"hourly"`
	Daily          []Daily    `json:"daily"`
	Alerts         []Alert    `json:"alerts"`
}

type Current struct {
	Dt         int64     `json:"dt"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Minutely struct {
	Dt            int64   `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type Hourly struct {
	Dt         int64     `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
}

type Daily struct {
	Dt        int64     `json:"dt"`
	Sunrise   int64     `json:"sunrise"`
	Sunset    int64     `json:"sunset"`
	Moonrise  int64     `json:"moonrise"`
	Moonset   int64     `json:"moonset"`
	MoonPhase float64   `json:"moon_phase"`
	Summary   string    `json:"summary"`
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   int       `json:"wind_deg"`
	WindGust  float64   `json:"wind_gust"`
	Weather   []Weather `json:"weather"`
	Clouds    int       `json:"clouds"`
	Pop       float64   `json:"pop"`
	Rain      float64   `json:"rain"`
	Uvi       float64   `json:"uvi"`
}

type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type Alert struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int64    `json:"start"`
	End         int64    `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
