package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	Start(router *gin.Engine, address string) error

	GetHealth(ctx *gin.Context)

	// Forecast
	GetWeatherForecast(ctx *gin.Context)
}
