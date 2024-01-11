package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	initialization "github.com/pafello/gocast/init"
	"github.com/pafello/gocast/internal/errors"
	"github.com/pafello/gocast/internal/geolocation"
	"github.com/pafello/gocast/internal/settings"
	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/weather"
)

func main() {
	initialization.InitEnv()
	var lat, lng float64

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
		ans = ans[:len(ans)-1]
		fmt.Println(ans)
		availableLocations, err := geolocation.GetGeolocationResults(ans)
		if err != nil {
			fmt.Println("Could not get locations results:", err)
		}
		for i, l := range availableLocations {

			fmt.Printf("%d: %s\n", i+1, l.Describe())
		}
		fmt.Print("Select your city (1/2/3): ")
		ans, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		option, err := strconv.Atoi(ans[:len(ans)-1])

		if err != nil {
			log.Fatal(err)
		}

		selectedResult := availableLocations[option-1]
		lat, lng = selectedResult.Lat, selectedResult.Lng

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
