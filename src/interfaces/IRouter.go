package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IRouter interface {
	GetRouter() *gin.Engine
	NewRouter(router *gin.Engine, server IServer) *gin.Engine
}
