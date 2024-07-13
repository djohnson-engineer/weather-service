package interfaces

import "weather-server/src/models"

type IWeatherService interface {
	GetForecast(lat, lon string) (*models.GetForecastResponse, error)
}
