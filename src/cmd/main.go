package main

import (
	"context"
	"weather-server/src/cmd/app"
	"weather-server/src/container"
	"weather-server/src/controllers"
	"weather-server/src/interfaces"
	"weather-server/src/logger"
)

var (
	routeObj interfaces.IRouter = &controllers.Router{}
	config   func()             = configureServices
)

// @securityDefinitions.apikey	Authorization
// @in							header
// @name						Authorization
// @description				Enter your OAuth Access Token here (prefix with 'bearer ')
func main() {
	// Dependency Injection
	config()

	// Initialize Router
	server := app.Instance.Server
	router := routeObj.GetRouter()
	newRouter := routeObj.NewRouter(router, server)

	// TODO Initialize and Configure Cache

	hostAndPort := app.Config().ServiceHost + ":" + app.Config().ServicePort
	logger.Log(logger.Info, "Starting weather-svc; URL is 'http://%s%s'", hostAndPort, app.Config().ServicePath)

	// Start Server
	if err := server.Start(newRouter, hostAndPort); err != nil {
		panic("Error starting server")
	}
}

// configureServices initializes the app services using dependency injection
func configureServices() {
	appContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	instance, err := container.CreateContainerInstance(appContext)
	if err != nil {
		panic(err) // Due to wire_gen created code, this currently cannot happen
	}

	app.Initialize(instance)
	app.Configure()
}
