package weather

import "github.com/pafello/gocast/internal/units"

type Forecast struct {
	List           []Weather `json:"list"`
	UnitSystemUsed units.UnitSystem
}

func (f *Forecast) Describe(cityName string) string {
	desc := ""
	for _, weather := range f.List {
		desc += weather.Describe(cityName)
	}
	return desc
}
