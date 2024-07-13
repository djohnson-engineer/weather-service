package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Configure(t *testing.T) {
	Configure()
	assert := assert.New(t)
	assert.Equal(nwsAPIURL, Config().NationalWeatherServiceApiURL)
	assert.Equal(userAgent, Config().UserAgent)
}
