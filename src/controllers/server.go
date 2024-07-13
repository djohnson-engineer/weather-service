package controllers

import (
	"weather-server/src/interfaces"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
}

func NewServer() *Server {
	return &Server{
		// TODO interface implementations here
	}
}

var _ interfaces.IServer = &Server{}

// Start runs the HTTP server on a specific address
func (*Server) Start(router *gin.Engine, address string) error {
	return router.Run(address)
}
