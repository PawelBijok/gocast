package weather

import (
	"fmt"

	"github.com/pafello/gocast/internal/units"
)

type Forecast struct {
	List           []Weather `json:"list"`
	UnitSystemUsed units.UnitSystem
}

func (f *Forecast) Describe(cityName string) string {
	desc := fmt.Sprintf("%s\n", cityName)

	for _, weather := range f.List {
		desc += weather.DescribeShort()
		desc += "\n"
	}
	return desc
}
