package openweather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type HTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type OpenWeather struct {
	APIKey   string
	Country  string
	ZipCode  string
	HTTPDoer HTTPDoer
}

type Probe struct {
	TempCelsius        float64
	HumidityPercentage float64
}

func New(APIKey string, country string, zipCode string) *OpenWeather {
	return &OpenWeather{
		APIKey:   APIKey,
		Country:  country,
		ZipCode:  zipCode,
		HTTPDoer: http.DefaultClient,
	}
}

type fivedayforecastResponse struct {
	HTTPCode string      `json:"cod"`
	Message  interface{} `json:"message"`
	Probes   []struct {
		Timestamp int64 `json:"dt"`
		Main      struct {
			Temp     float64 `json:"temp"`
			Humidity float64 `json:"humidity"`
		} `json:"main"`
	} `json:"list"`
}

func (o *OpenWeather) Get(ctx context.Context, from time.Time, to time.Time) ([]Probe, error) {
	request, _ := http.NewRequestWithContext(ctx, "GET", o.fiveDayForecastURL(), nil)
	response, err := o.HTTPDoer.Do(request)
	if err != nil {
		return nil, fmt.Errorf("can't request openweather API: %w", err)
	}

	var allprobes fivedayforecastResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&allprobes); err != nil {
		return nil, fmt.Errorf("can't decode openweather response: %w", err)
	}

	if allprobes.HTTPCode != "200" {
		return nil, fmt.Errorf("can't get valid openweather response: %#+v", allprobes)
	}

	var nightprobes []Probe
	for _, probe := range allprobes.Probes {
		probeTime := time.Unix(probe.Timestamp, 0)
		if probeTime.After(from) && probeTime.Before(to) {
			nightprobes = append(nightprobes, Probe{
				TempCelsius:        probe.Main.Temp,
				HumidityPercentage: probe.Main.Humidity,
			})
		}
	}

	return nightprobes, nil
}

func (o *OpenWeather) fiveDayForecastURL() string {
	forecast, _ := url.Parse("https://api.openweathermap.org/data/2.5/forecast")

	querystring := forecast.Query()
	querystring.Set("appid", o.APIKey)
	querystring.Set("zip", fmt.Sprintf("%s,%s", o.ZipCode, o.Country))
	querystring.Set("units", "metric")
	forecast.RawQuery = querystring.Encode()

	return forecast.String()
}
