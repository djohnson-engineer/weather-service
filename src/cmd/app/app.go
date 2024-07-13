package app

import "weather-server/src/interfaces"

var Instance *Container

type Container struct {
	Server interfaces.IServer
}

// Dependency Injection
func Initialize(container *Container) {
	Instance = container
}
