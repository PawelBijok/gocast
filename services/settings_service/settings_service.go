package settings_service

import (
	"encoding/json"
	"errors"
	"os"

	e "github.com/pafello/gocast/services/settings_service/errors"
	m "github.com/pafello/gocast/services/settings_service/models"
)

type SettingsService struct {
	JsonFilePath string
}

func (s *SettingsService) GetUserSettings() (m.UserSettings, error) {

	fileBytes, err := os.ReadFile(s.JsonFilePath)
	if err != nil {
		return m.UserSettings{}, errors.New(e.FileDoesNotExists)
	}
	us := m.UserSettings{}
	err = json.Unmarshal(fileBytes, &us)

	if err != nil {
		return m.UserSettings{}, errors.New(e.CantUnmarshalJson)
	}
	return us, nil
}
