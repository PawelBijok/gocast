package settings

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pafello/gocast/internal/geolocation"
	"github.com/pafello/gocast/internal/units"
)

func InterviewUser() (UserSettings, error) {

	printHeader("Location")
	fmt.Print("Specify you location (ex. Helsinki): ")
	var ans string

	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ans = ans[:len(ans)-1]
	availableLocations, err := geolocation.GetGeolocationResults(ans)
	if err != nil {
		log.Fatal("Could not get locations results:", err)
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
	printHeader("Unit system")
	fmt.Println("1. Metric")
	fmt.Println("2. Imperial")
	fmt.Print("Select your unit system: ")
	ans, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	var unitSys units.UnitSystem
	if option == 1 {
		unitSys = units.Metric
	} else {
		unitSys = units.Imperial
	}

	userSettings := UserSettings{
		Location: selectedResult,
		UnitSys:  unitSys,
	}

	return userSettings, nil
}

func printHeader(title string) {
	fmt.Println("")
	fmt.Println(title)
	underscore := ""
	for i := 0; i < len(title); i++ {

		underscore += "‾"
	}
	fmt.Println(underscore)
}
