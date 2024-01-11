package weather

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pafello/gocast/config"
	"github.com/pafello/gocast/internal/units"
)

type WeatherService interface {
	GetWeather(lat float64, lon float64) (Weather, error)
}

func createWeatherUrl(lat float64, lng float64, unitSystem units.UnitSystem) string {
	latString, lonString := strconv.FormatFloat(lat, 'f', 6, 64), strconv.FormatFloat(lng, 'f', 6, 64)
	apiKey := os.Getenv(config.WeatherApiKey)
	params := map[string]string{
		"lat":   latString,
		"lon":   lonString,
		"units": string(unitSystem),
		"appId": apiKey,
	}
	baseUrl := "https://api.openweathermap.org/data/2.5/weather"
	url := baseUrl + "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}
	return url
}

func GetWeather(lat float64, lng float64, unitSystem units.UnitSystem) (Weather, error) {

	url := createWeatherUrl(lat, lng, unitSystem)

	res, err := http.Get(url)
	if err != nil {
		return Weather{}, err
	}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return Weather{}, err
	}

	weather := Weather{}
	err = json.Unmarshal([]byte(jsonData), &weather)
	if err != nil {
		return Weather{}, err
	}
	weather.UnitSystemUsed = unitSystem
	return weather, nil
}
