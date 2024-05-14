package app_mode

import "errors"

type AppMode string

const (
	Help              AppMode = "help"
	ChangePreferences AppMode = "prefs"
	Current           AppMode = "current"
	Forecast          AppMode = "forecast"
)

func AppModeFromString(s string) (AppMode, error) {
	switch s {
	case string(Help):
		return Help, nil
	case string(ChangePreferences):
		return ChangePreferences, nil
	case string(Current):
		return Current, nil
	case string(Forecast):
		return Forecast, nil
	}
	return "", errors.New("No matching elements")

}
