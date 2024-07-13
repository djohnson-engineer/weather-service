package interfaces

import "weather-server/src/domain"

// Defines the contract for fetching weather data from the third-party API
type IForecast interface {
	GetForecastURL(lat, lon string) (*string, error)
	GetWeatherData(forecastURL string) (*domain.WeatherForecast, error)
}
