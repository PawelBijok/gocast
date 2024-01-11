package settings

import (
	"encoding/json"
	"errors"
	"os"

	e "github.com/pafello/gocast/internal/errors"
)

type SettingsService struct {
	JsonFilePath string
}

func (s *SettingsService) GetUserSettings() (UserSettings, error) {

	fileBytes, err := os.ReadFile(s.JsonFilePath)
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
