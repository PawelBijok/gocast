package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	initialization "github.com/pafello/gocast/init"
	"github.com/pafello/gocast/internal/errors"
	"github.com/pafello/gocast/internal/settings"
	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/weather"
)

func main() {
	initialization.InitEnv()
	lat, lng := 49.797417, 18.790270

	settingsService := settings.SettingsService{JsonFilePath: "config.json"}

	userSettings, err := settingsService.GetUserSettings()

	if err.Error() == errors.FileDoesNotExists {
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

	weather, err := weather.GetWeather(lat, lng, units.Metric)

	if err != nil {
		fmt.Println("Could not read the weather:", err)
		return
	}

	fmt.Println(weather.Describe())
	fmt.Println(weather.DescribeDetails())
}
