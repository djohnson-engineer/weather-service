package managers

import (
	"weather-server/src/datasource"
	"weather-server/src/interfaces"
	"weather-server/src/models"
	"weather-server/src/translation"
)

type DefaultWeatherService struct {
	api         interfaces.IForecast
	categorizer interfaces.ITemperatureCharacterizer
}

// TODO consider factory pattern or other option for different Forecast managers
func DefaultWeatherForecaster() *DefaultWeatherService {
	return &DefaultWeatherService{
		api:         &datasource.WeatherGovAPI{},
		categorizer: &translation.DefaultTemperatureCategorizer{},
	}
}

func (s *DefaultWeatherService) GetForecast(lat, lon string) (*models.GetForecastResponse, error) {
	forecastURL, err := s.api.GetForecastURL(lat, lon)
	if err != nil || forecastURL == nil {
		return nil, err
	}

	data, err := s.api.GetWeatherForecast(*forecastURL)
	if err != nil {
		return nil, err
	}

	temperatureType := s.categorizer.CharacterizeTemperature(data.Temperature)

	return &models.GetForecastResponse{
		ShortForecast:   data.ShortForecast,
		TemperatureType: temperatureType,
	}, nil
}
