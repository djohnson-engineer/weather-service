package interfaces

import "weather-server/src/domain"

// Defines the contract for the weather forecast service
type IWeatherService interface {
	GetForecast(lat, lon string) (*domain.WeatherResponse, error)
}
