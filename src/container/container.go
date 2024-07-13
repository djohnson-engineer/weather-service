//go:build wireinject
// +build wireinject

package container

import (
	"context"
	"weather-server/src/cmd/app"
	"weather-server/src/controllers"
	"weather-server/src/interfaces"

	"github.com/google/wire"
)

var server = wire.NewSet(
	controllers.NewServer,
	wire.Bind(new(interfaces.IServer), new(*controllers.Server)),
)

// Example code for DI Implementation
func CreateContainerInstance(ctx context.Context) (*app.Container, error) {
	panic(wire.Build(
		server,
		wire.Struct(new(app.Container), "*"),
	))
}
