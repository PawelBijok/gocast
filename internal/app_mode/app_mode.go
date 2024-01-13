package app_mode

import "errors"

type AppMode string

const (
	Help              AppMode = "help"
	ChangePreferences AppMode = "prefs"
	Normal            AppMode = "normal"
)

func AppModeFromString(s string) (AppMode, error) {
	switch s {
	case string(Help):
		return Help, nil
	case string(ChangePreferences):
		return ChangePreferences, nil
	case string(Normal):
		return Normal, nil
	}
	return "", errors.New("No matching elements")

}
