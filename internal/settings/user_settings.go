package settings

import (
	"github.com/pafello/gocast/internal/geolocation"
	"github.com/pafello/gocast/internal/units"
)

type UserSettings struct {
	UnitSys  units.UnitSystem              `json:"unitSys"`
	Location geolocation.GeolocationResult `json:"location"`
}
