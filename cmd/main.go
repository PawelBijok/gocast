package main

import (
	"fmt"
	"log"
	"os"

	initialization "github.com/pafello/gocast/init"
	"github.com/pafello/gocast/internal/app_mode"
	"github.com/pafello/gocast/internal/help"
	"github.com/pafello/gocast/internal/settings"
	"github.com/pafello/gocast/internal/weather"
)

type AppMode string

func main() {

	initialization.InitEnv()
	args := os.Args
	var mode app_mode.AppMode
	if len(args) > 1 {
		arg := args[1]
		var err error
		mode, err = app_mode.AppModeFromString(arg)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		mode = app_mode.Normal
	}

	if mode == app_mode.Help {

		help.ShowHelp()
		return
	}
	userSettings, err := settings.GetUserSettings()

	if err != nil || mode == app_mode.ChangePreferences {
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

	if mode == app_mode.Forecast {
		forecast, err := weather.GetForecast(userSettings.Location.Lat, userSettings.Location.Lng, userSettings.UnitSys)
		if err != nil {
			fmt.Println("Could not read forecast:", err)
			return
		}
		groups := forecast.GroupWeatherByDay()
		groups.DescribeDaily()

	} else {

		weather, err := weather.GetWeather(userSettings.Location.Lat, userSettings.Location.Lng, userSettings.UnitSys)

		if err != nil {
			fmt.Println("Could not read the weather:", err)
			return
		}

		fmt.Println(weather.Describe(userSettings.Location.Name))
		fmt.Println(weather.DescribeDetails())
	}
}
