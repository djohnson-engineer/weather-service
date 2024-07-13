package controllers

import (
	"testing"
	"weather-server/src/mocks"
)

type ServerMocks struct {
	weatherService *mocks.IWeatherService
}

func NewTestServer(t *testing.T) (Server, ServerMocks) {
	mox := ServerMocks{
		weatherService: mocks.NewIWeatherService(t),
	}

	return *NewServer(), mox
}
