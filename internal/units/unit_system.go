package units

type UnitSystem string

const (
	Metric   UnitSystem = "metric"
	Imperial UnitSystem = "imperial"
)

func (u UnitSystem) GetTempUnit() string {
	if u == Metric {
		return "°C"
	} else {
		return "°F"
	}
}
func (u UnitSystem) GetSpeedUnit() string {
	if u == Metric {
		return "km/h"
	} else {
		return "mph"
	}
}
func (u UnitSystem) GetPressureUnit() string {
	if u == Metric {
		return "hPa"
	} else {
		return "PSI"
	}
}
