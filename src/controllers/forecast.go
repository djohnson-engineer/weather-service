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
//	@Security	Authorization
//	@Accept		json
//	@Produce	json
//	@Param		latitude	path		string	true	"Latitude"
//	@Param		longitude	path		string	true	"Longitude"
//	@Success	200		{object}	domain.WeatherForecast
//	@Failure	400		{object}	Response
//	@Failure	404		{object}	Response
//	@Router		/forecast?latitude={latitude}&longitude={longitude} [get]
func (s *Server) GetWeatherForecast(c *gin.Context) {
	// latitude := c.Param("latitude")
	// longitude := c.Param("longitude")

	lat := c.Query("latitude")
	lon := c.Query("longitude")

	if lat == "" || lon == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Latitude and longitude are required"})
		return
	}

	forecaster := managers.DefaultWeatherForecaster()

	//TODO need to evaluate various status codes depending on the error/issue
	response, err := forecaster.GetForecast(lat, lon)
	if err != nil {
		logger.LogError(c, "error getting forecast: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger.LogInfo(c, "successfully returned forecast for %s,%s: '%s'", lat, lon, response)
	c.JSON(http.StatusOK, response)
}
