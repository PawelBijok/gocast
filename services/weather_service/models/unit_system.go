package models

type UnitSystem string

const (
	Metric   UnitSystem = "metric"
	Imperial UnitSystem = "imperial"
)

func (u UnitSystem) getTempUnit() string {
	if u == Metric {
		return "°C"
	} else {
		return "°F"
	}
}
func (u UnitSystem) getSpeedUnit() string {
	if u == Metric {
		return "km/h"
	} else {
		return "mph"
	}
}
func (u UnitSystem) getPressureUnit() string {
	if u == Metric {
		return "hPa"
	} else {
		return "PSI"
	}
}
