package managers

import (
	"weather-server/src/domain"
	"weather-server/src/interfaces"
	"weather-server/src/thirdparty"
	"weather-server/src/translation"
)

type DefaultWeatherService struct {
	api         interfaces.IForecast
	categorizer interfaces.ITemperatureCharacterizer
}

// TODO consider factory pattern or other option for different Forecast managers
func DefaultWeatherForecaster() *DefaultWeatherService {
	return &DefaultWeatherService{
		api:         &thirdparty.WeatherGovAPI{},
		categorizer: &translation.DefaultTemperatureCategorizer{},
	}
}

func (s *DefaultWeatherService) GetForecast(lat, lon string) (*domain.WeatherResponse, error) {
	forecastURL, err := s.api.GetForecastURL(lat, lon)
	if err != nil || forecastURL == nil {
		return nil, err
	}

	data, err := s.api.GetWeatherData(*forecastURL)
	if err != nil {
		return nil, err
	}

	temperatureType := s.categorizer.CharacterizeTemperature(data.Temperature)

	return &domain.WeatherResponse{
		ShortForecast:   data.ShortForecast,
		TemperatureType: temperatureType,
	}, nil
}
