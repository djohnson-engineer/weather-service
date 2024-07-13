package thirdparty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-server/src/cmd/app"
	"weather-server/src/domain"
)

func setUserAgent(req *http.Request) {
	req.Header.Set("User-Agent", app.Config().UserAgent)
}

type WeatherGovAPI struct{}

func (api *WeatherGovAPI) GetForecastURL(lat, lon string) (*string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(app.Config().NationalWeatherServiceApiURL, lat, lon), nil)
	if err != nil {
		return nil, err
	}
	setUserAgent(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	//TODO if errors:
	// headers contained "X-Request-Id" and "X-Edge-Request-Id" that you can use to discuss with third-party

	/*
		Source: https://www.weather.gov/documentation/services-web-api
		Note: Applications may cache the grid for a location to improve latency and reduce the additional lookup request; however, it is important to note that while it generally does not occur often, the gridX and gridY values (and even the office) for a given coordinate may occasionally change. For this reason, it is necessary to check back to the /points endpoint periodically for the latest office/grid mapping
	*/

	var pointsResponse natWeatherServicePointsResponse
	if err := json.NewDecoder(resp.Body).Decode(&pointsResponse); err != nil {
		return nil, err
	}

	return &pointsResponse.Properties.Forecast, nil
}

type natWeatherServicePointsResponse struct {
	Properties struct {
		Forecast string `json:"forecast"`
	} `json:"properties"`
}

type natWeatherForecastResponse struct {
	Properties struct {
		Periods []struct {
			Name          string `json:"name"`
			Temperature   int    `json:"temperature"`
			ShortForecast string `json:"shortForecast"`
		} `json:"periods"`
	} `json:"properties"`
}

func (api *WeatherGovAPI) GetWeatherData(forecastURL string) (*domain.WeatherForecast, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", forecastURL, nil)
	if err != nil {
		return nil, err
	}
	setUserAgent(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get weather data: %s", resp.Status)
	}

	var response natWeatherForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	if len(response.Properties.Periods) == 0 {
		return nil, fmt.Errorf("no forecast data available")
	}

	todayForecast := response.Properties.Periods[0]

	return &domain.WeatherForecast{
		Temperature:   todayForecast.Temperature,
		ShortForecast: todayForecast.ShortForecast,
	}, nil
}
