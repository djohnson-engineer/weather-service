package domain

// WeatherForecast represents raw weather data from the datasource
type WeatherForecast struct {
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"short_forecast"`
}
