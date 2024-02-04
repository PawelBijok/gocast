package settings

import (
	"encoding/json"
	"errors"
	"os"

	e "github.com/pafello/gocast/internal/errors"
)

const settingsFilePath = "/go_cast_settings.json"
const settingsDirectoryPath = "/.config/gocast"

func getConfigDirectory() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dirname + settingsDirectoryPath, nil
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
		return errors.New(e.CantMarshalJson)
	}

	filePath, err := getConfigDirectory()
	if err != nil {
		return errors.New(e.CantCreateConfigDirectory)
	}

	os.MkdirAll(filePath, 0755)
	filePath, err = getConfigFilePath()
	if err != nil {
		return errors.New(e.CantCreateConfigDirectory)
	}

	file, err := os.Create(filePath)

	_, err = file.Write(bytes)

	if err != nil {
		return errors.New(e.CanSaveToFile)
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
