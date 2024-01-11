package playground

//interfaces explained
import (
	"fmt"
	"strings"
)

type WeatherService interface {
	GetWeather() string
}

type OpenWeather struct {
	OpenWeatherKey string
	ApiVersion     string
}
type GoogleWeather struct {
	GoogleVersion string
}

func (g GoogleWeather) GetWeather() string {

	return "sunny google"
}
func (g OpenWeather) GetWeather() string {

	return "sunny google"
}

func testWeatherService(service WeatherService) {
	if strings.Contains(service.GetWeather(), "google") {
		fmt.Println("its goolge")
	} else {

		fmt.Println("its something else")
	}
}

func main() {
	test := GoogleWeather{GoogleVersion: "v1.1.1"}
	testWeatherService(test)
}
