package app_mode

import "errors"

type AppMode string

const (
	Help              AppMode = "help"
	ChangePreferences AppMode = "prefs"
	Normal            AppMode = "normal"
	Forecast          AppMode = "forecast"
)

func AppModeFromString(s string) (AppMode, error) {
	switch s {
	case string(Help):
		return Help, nil
	case string(ChangePreferences):
		return ChangePreferences, nil
	case string(Normal):
		return Normal, nil
	case string(Forecast):
		return Forecast, nil
	}
	return "", errors.New("No matching elements")

}
