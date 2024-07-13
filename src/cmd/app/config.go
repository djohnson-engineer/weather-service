package app

import (
	"os"
	"strings"
)

var (
	config *AppConfig
)

const (
	nwsAPIURL = "https://api.weather.gov/points/%s,%s"
	userAgent = "MyWeatherForecastApp/1.0 (teamemail@company.com)"

	defaultHost        = "0.0.0.0"
	defaultPort        = "8950"
	defaultServicePath = "/api/weather/v1/"
)

type AppConfig struct {
	ServiceHost string
	ServicePort string
	ServicePath string

	NationalWeatherServiceApiURL string
	UserAgent                    string
}

func Config() AppConfig {
	if config == nil {
		Configure()
	}

	return *config
}

func Configure() {
	env := getEnv()

	config = &AppConfig{
		ServiceHost:                  getEnvVal(env, "HOST", defaultHost),
		ServicePort:                  getEnvVal(env, "PORT", defaultPort),
		NationalWeatherServiceApiURL: nwsAPIURL,
		UserAgent:                    userAgent,
	}
}

func getEnv() map[string]string {
	env := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair == nil {
			return env
		}

		env[pair[0]] = pair[1]
	}
	return env
}

func getEnvVal(env map[string]string, key string, defaultVal string) string {
	val := env[key]
	if len(val) == 0 {
		val = defaultVal
	}
	return val
}
