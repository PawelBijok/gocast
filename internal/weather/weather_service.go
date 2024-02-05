package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pafello/gocast/config"
	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/utils"
)

func crateOpenWeatherUrl(lat float64, lng float64, unitSystem units.UnitSystem, apiCallType apiCallType) string {
	latString, lonString := strconv.FormatFloat(lat, 'f', 6, 64), strconv.FormatFloat(lng, 'f', 6, 64)
	apiKey := os.Getenv(config.OpenWeatherApiKey)
	params := map[string]string{
		"lat":   latString,
		"lon":   lonString,
		"units": string(unitSystem),
		"appId": apiKey,
	}
	baseUrl := fmt.Sprintf("https://api.openweathermap.org/data/2.5/%s", string(apiCallType))
	return utils.GenerateUrl(baseUrl, params)
}

func GetWeather(lat float64, lng float64, unitSystem units.UnitSystem) (Weather, error) {

	url := crateOpenWeatherUrl(lat, lng, unitSystem, weather)

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

func GetForecast(lat float64, lng float64, unitSystem units.UnitSystem) (Forecast, error) {
	url := crateOpenWeatherUrl(lat, lng, unitSystem, forecast)

	res, err := http.Get(url)
	if err != nil {
		return Forecast{}, err
	}

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return Forecast{}, err
	}
	forecast := Forecast{}
	err = json.Unmarshal([]byte(jsonData), &forecast)
	if err != nil {
		return Forecast{}, err
	}
	forecast.UnitSystemUsed = unitSystem
	return forecast, nil
}
