package controllers

import (
	"weather-server/src/interfaces"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func (*Router) GetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func (*Router) NewRouter(r *gin.Engine, server interfaces.IServer) *gin.Engine {
	baseUrl := "/api/weather/v1/"
	router := r.Group(baseUrl)

	//TODO auth layer

	// Health API
	router.GET("/health", server.GetHealth)

	// Forecast API
	router.GET("/forecast", server.GetWeatherForecast)

	// Docs API
	// TODO need a simple documentation endpoint (e.g. swagger)
	// router.GET("/docs/*any", <insert-doc-endpoint>)

	return r
}
