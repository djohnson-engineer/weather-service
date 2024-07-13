package models

// GetForecastResponse represents the response from the weather service
type GetForecastResponse struct {
	ShortForecast   string `json:"short_forecast"`
	TemperatureType string `json:"temperature_type"`
}
