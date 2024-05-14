package help

import "fmt"

func ShowHelp() {

	fmt.Println(" ")
	fmt.Println("Available options: current | forecast | help | prefs")
	fmt.Println(" ")
	fmt.Println("current: (optional) runs the app in default mode")
	fmt.Println("forecast: shows the forecast forn next 4 hours")
	fmt.Println("help: display app help")
	fmt.Println("prefs: change current preferences")
}
