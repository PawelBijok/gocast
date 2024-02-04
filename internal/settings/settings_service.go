package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	e "github.com/pafello/gocast/internal/errors"
)

const settingsFilePath = "/go_cast_settings.json"

func getConfigDirectory() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname + "/.config/gocast", nil
}
func getConfigFilePath() (string, error) {
	dirname, err := getConfigDirectory()
	if err != nil {
		return "", err
	}
	return dirname + settingsFilePath, nil
}

func SaveUserSettings(settings UserSettings) error {
	bytes, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	filePath, err := getConfigDirectory()
	if err != nil {
		fmt.Println("Could not get config directory")
		return err
	}

	os.MkdirAll(filePath, 0755)
	filePath, err = getConfigFilePath()
	if err != nil {
		fmt.Println("Could not get config file")
		return err
	}

	file, err := os.Create(filePath)

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println("Could not write to file")
		fmt.Println(err)
		return err
	}
	return nil

}
func GetUserSettings() (UserSettings, error) {

	filePath, err := getConfigFilePath()
	if err != nil {

		return UserSettings{}, errors.New(e.FileDoesNotExists)
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return UserSettings{}, errors.New(e.FileDoesNotExists)
	}
	us := UserSettings{}
	err = json.Unmarshal(fileBytes, &us)

	if err != nil {
		return UserSettings{}, errors.New(e.CantUnmarshalJson)
	}
	return us, nil
}
