package main

import (
	"fmt"
	"net/http"
	"os"
	"weather-server/src/domain"
	"weather-server/src/interfaces"
	"weather-server/src/thirdparty"
	"weather-server/src/translation"

	"github.com/gin-gonic/gin"
)

type DefaultWeatherService struct {
	api         interfaces.ForecastAPI
	categorizer interfaces.TemperatureCharacterizer
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

func weatherHandler(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")

	if lat == "" || lon == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
		return
	}

	service := &DefaultWeatherService{
		api:         &thirdparty.WeatherGovAPI{},
		categorizer: &translation.DefaultTemperatureCategorizer{},
	}

	response, err := service.GetForecast(lat, lon)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Forecast: %s, Temperature Type: %s\n", response.ShortForecast, response.TemperatureType)
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.GET("/weather", weatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
