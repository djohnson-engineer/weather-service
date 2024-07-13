package controllers

import (
	"weather-server/src/interfaces"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests.
type Server struct {
	// managers interfaces.IManagerProvider
}

func NewServer(
// managers interfaces.IManagerProvider,
) *Server {
	return &Server{
		// managers: managers,
	}
}

var _ interfaces.IServer = &Server{}

// Start runs the HTTP server on a specific address.
func (*Server) Start(router *gin.Engine, address string) error {
	return router.Run(address)
}
