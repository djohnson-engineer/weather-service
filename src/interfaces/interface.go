package interfaces

import "weather-server/src/domain"

// WeatherService defines the contract for the weather forecast service
type WeatherService interface {
	GetForecast(lat, lon string) (*domain.WeatherResponse, error)
}

// ForecastAPI defines the contract for fetching weather data from the third-party API
type ForecastAPI interface {
	GetForecastURL(lat, lon string) (*string, error)
	GetWeatherData(forecastURL string) (*domain.WeatherData, error)
}

// TemperatureCharacterizer defines the contract for characterizing temperatures.
type TemperatureCharacterizer interface {
	CharacterizeTemperature(temp int) string
}
