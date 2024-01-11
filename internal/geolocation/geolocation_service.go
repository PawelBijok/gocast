package geolocation

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/pafello/gocast/config"
	"github.com/pafello/gocast/internal/utils"
)

func createGeolocationUrl(cityName string, limit int) string {
	apiKey := os.Getenv(config.OpenWeatherApiKey)
	params := map[string]string{
		"q":     cityName,
		"limit": strconv.Itoa(limit),
		"appId": apiKey,
	}
	baseUrl := "http://api.openweathermap.org/geo/1.0/direct"
	return utils.GenerateUrl(baseUrl, params)
}

func GetGeolocationResults(cityName string) ([]GeolocationResult, error) {
	url := createGeolocationUrl(cityName, 3)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {

		return nil, err
	}
	var data []GeolocationResult
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
