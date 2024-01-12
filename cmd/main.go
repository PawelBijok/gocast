package main

import (
	"fmt"
	"log"

	initialization "github.com/pafello/gocast/init"
	"github.com/pafello/gocast/internal/settings"
	"github.com/pafello/gocast/internal/weather"
)

func main() {
	initialization.InitEnv()

	userSettings, err := settings.GetUserSettings()

	if err != nil {
		fmt.Println("Welcome to GO Cast!")

		userSettings, err = settings.InterviewUser()
		if err != nil {
			log.Fatal("Could not get your preferences...")
		}

		err = settings.SaveUserSettings(userSettings)
		if err != nil {
			log.Fatal("Could not save the preferences...")
		}
		fmt.Println("Your preferences have been saved :)")
		fmt.Println("")
	}

	weather, err := weather.GetWeather(userSettings.Location.Lat, userSettings.Location.Lng, userSettings.UnitSys)

	if err != nil {
		fmt.Println("Could not read the weather:", err)
		return
	}

	fmt.Println(weather.Describe(userSettings.Location.Name))
	fmt.Println(weather.DescribeDetails())
}
