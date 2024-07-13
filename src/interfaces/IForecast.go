package interfaces

import "weather-server/src/domain"

// Defines the contract for fetching weather data from various datasources
type IForecast interface {
	GetForecastURL(lat, lon string) (*string, error)
	GetWeatherForecast(forecastURL string) (*domain.WeatherForecast, error)
}
