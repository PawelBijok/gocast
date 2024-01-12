package settings

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	e "github.com/pafello/gocast/internal/errors"
)

const settingsFilePath = "./go_cast_settings.json"

func SaveUserSettings(settings UserSettings) error {
	bytes, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	err = os.WriteFile(settingsFilePath, bytes, 0666)
	if err != nil {
		fmt.Println("Could not write to file")
		return err
	}
	return nil

}
func GetUserSettings() (UserSettings, error) {

	fileBytes, err := os.ReadFile(settingsFilePath)
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
