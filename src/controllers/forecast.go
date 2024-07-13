package controllers

import (
	"net/http"
	"weather-server/src/logger"
	"weather-server/src/managers"

	"github.com/gin-gonic/gin"
)

// GetWeatherForecast godoc
//
//	@Summary	Get weather forecaset
//	@Schemes
//	@Tags		forecast
//	@Accept		json
//	@Produce	json
//	@Param		latitude	path		string	true	"Latitude"
//	@Param		longitude	path		string	true	"Longitude"
//	@Success	200		{object}	domain.WeatherForecast
//	@Failure	400		{object}	Response
//	@Failure	404		{object}	Response
//	@Router		/forecast?latitude={latitude}&longitude={longitude} [get]
func (s *Server) GetWeatherForecast(c *gin.Context) {
	lat := c.Query("latitude")
	lon := c.Query("longitude")

	//TODO common validation library
	if len(lat) == 0 || len(lon) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
		return
	}

	forecaster := managers.DefaultWeatherForecaster()

	//TODO need to evaluate various status codes depending on the error/issue
	response, err := forecaster.GetForecast(lat, lon)
	if err != nil {
		logger.Log(logger.Error, "error getting forecast: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if response == nil {
		logger.Log(logger.Warning, "no forecast found for Latitude: %s, Longitude: %s.", lat, lon)
		c.JSON(http.StatusNoContent, nil)
		return
	}

	logger.Log(logger.Info, "successfully returned forecast for Latitude: %s, Longitude: %s: '%s'", lat, lon, *response)
	c.JSON(http.StatusOK, response)
}
