package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	i "github.com/pafello/gocast/init"
	ss "github.com/pafello/gocast/services/settings_service"
	sse "github.com/pafello/gocast/services/settings_service/errors"
	ws "github.com/pafello/gocast/services/weather_service"
	wsm "github.com/pafello/gocast/services/weather_service/models"
)

func main() {
	i.InitEnv()
	lat, lng := 49.797417, 18.790270

	settingsService := ss.SettingsService{JsonFilePath: "config.json"}

	userSettings, err := settingsService.GetUserSettings()

	if err.Error() == sse.FileDoesNotExists {
		fmt.Println("Welcome to GO Cast!")
		fmt.Print("Specify you location (ex. Los Angeles): ")
		var ans string

		reader := bufio.NewReader(os.Stdin)
		ans, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(ans)
	}

	fmt.Println(userSettings)

	weather, err := ws.GetWeather(lat, lng, wsm.Metric)

	if err != nil {
		fmt.Println("Could not read the weather:", err)
		return
	}

	fmt.Println(weather.Describe())
	fmt.Println(weather.DescribeDetails())
}
