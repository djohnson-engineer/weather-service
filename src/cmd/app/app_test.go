package app_test

import (
	"testing"
	"weather-server/src/cmd/app"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	var container *app.Container

	app.Initialize(container)
	assert.Equal(t, container, app.Instance)
}
