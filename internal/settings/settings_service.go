package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	e "github.com/pafello/gocast/internal/errors"
)

const settingsFilePath = "/go_cast_settings.json"

func getConfigFile() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname + "/.config" + settingsFilePath, nil
}

func SaveUserSettings(settings UserSettings) error {
	bytes, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	filePath, err := getConfigFile()
	if err != nil {

		fmt.Println("Could not get config file")
		return err
	}
	err = os.WriteFile(filePath, bytes, 0666)
	if err != nil {
		fmt.Println("Could not write to file")
		return err
	}
	return nil

}
func GetUserSettings() (UserSettings, error) {

	filePath, err := getConfigFile()
	if err != nil {

		fmt.Println("Could not get config file")
		return UserSettings{}, err
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could not get file")
		return UserSettings{}, errors.New(e.FileDoesNotExists)
	}
	us := UserSettings{}
	err = json.Unmarshal(fileBytes, &us)

	if err != nil {
		fmt.Println("Could not get file")
		return UserSettings{}, errors.New(e.CantUnmarshalJson)
	}
	return us, nil
}
