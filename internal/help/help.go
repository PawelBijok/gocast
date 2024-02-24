package help

import "fmt"

func ShowHelp() {

	fmt.Println(" ")
	fmt.Println("Available options: help | prefs | normal")
	fmt.Println(" ")
	fmt.Println("help: display app help")
	fmt.Println("prefs: change current preferences")
	fmt.Println("forecast: shows the forecast forn next 4 hours")
	fmt.Println("normal: (optional) runs the app in default mode")
}
