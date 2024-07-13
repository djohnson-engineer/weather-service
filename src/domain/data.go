package domain

// WeatherResponse represents the response from the weather service.
type WeatherResponse struct {
	ShortForecast   string `json:"short_forecast"`
	TemperatureType string `json:"temperature_type"`
}

// WeatherData represents the raw weather data from an external API.
type WeatherData struct {
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"short_forecast"`
}
