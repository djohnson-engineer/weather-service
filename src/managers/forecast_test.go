package managers

import (
	"errors"
	"testing"
	"weather-server/src/domain"
	"weather-server/src/mocks"
	"weather-server/src/models"
	"weather-server/src/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TODO if more time, consider refactoring tests to avoid having to say whether you expect an assertion or not
func createMockForecast(t *testing.T, url string, weatherData *domain.WeatherForecast, urlErr error, dataErr error, expectURL bool, expectData bool) *mocks.IForecast {
	mockForecast := mocks.NewIForecast(t)
	if expectURL {
		mockForecast.On("GetForecastURL", mock.Anything, mock.Anything).Return(func(lat, lon string) *string {
			if urlErr != nil {
				return nil
			}
			return &url
		}, urlErr)
	}
	if expectData {
		mockForecast.On("GetWeatherForecast", url).Return(weatherData, dataErr)
	}
	return mockForecast
}

func createMockCategorizer(t *testing.T, temperatureType string) *mocks.ITemperatureCharacterizer {
	mockCategorizer := mocks.NewITemperatureCharacterizer(t)
	mockCategorizer.On("CharacterizeTemperature", mock.Anything).Return(temperatureType)
	return mockCategorizer
}

func TestGetForecast(t *testing.T) {
	tests := []struct {
		name            string
		lat             string
		lon             string
		mockForecast    func() *mocks.IForecast
		mockCategorizer func() *mocks.ITemperatureCharacterizer
		expected        *models.GetForecastResponse
		expectErr       *string
	}{
		{
			name: "successful forecast retrieval",
			lat:  "40.7128",
			lon:  "-74.0060",
			mockForecast: func() *mocks.IForecast {
				url := "http://example.com/forecast"
				weatherData := &domain.WeatherForecast{
					Temperature:   75,
					ShortForecast: "Sunny",
				}
				return createMockForecast(t, url, weatherData, nil, nil, true, true)
			},
			mockCategorizer: func() *mocks.ITemperatureCharacterizer {
				return createMockCategorizer(t, "Warm")
			},
			expected: &models.GetForecastResponse{
				ShortForecast:   "Sunny",
				TemperatureType: "Warm",
			},
			expectErr: nil,
		},
		{
			name: "error fetching forecast URL",
			lat:  "40.7128",
			lon:  "-74.0060",
			mockForecast: func() *mocks.IForecast {
				return createMockForecast(t, "", nil, errors.New("failed to get forecast URL"), nil, true, false)
			},
			mockCategorizer: func() *mocks.ITemperatureCharacterizer {
				return nil
			},
			expected:  nil,
			expectErr: utils.StringPtr("failed to get forecast URL"),
		},
		{
			name: "error fetching weather forecast data",
			lat:  "40.7128",
			lon:  "-74.0060",
			mockForecast: func() *mocks.IForecast {
				url := "http://example.com/forecast"
				return createMockForecast(t, url, nil, nil, errors.New("failed to get weather forecast data"), true, true)
			},
			mockCategorizer: func() *mocks.ITemperatureCharacterizer {
				return nil
			},
			expected:  nil,
			expectErr: utils.StringPtr("failed to get weather forecast data"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &DefaultWeatherService{
				api:         tt.mockForecast(),
				categorizer: nil,
			}
			if tt.mockCategorizer != nil {
				service.categorizer = tt.mockCategorizer()
			}

			result, err := service.GetForecast(tt.lat, tt.lon)
			if tt.expectErr != nil {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), *tt.expectErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
