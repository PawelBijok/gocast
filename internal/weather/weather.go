package weather

import (
	"fmt"

	"github.com/pafello/gocast/internal/units"
)

type weatherCore struct {
	Temp          float32 `json:"temp"`
	TempFeelsLike float32 `json:"feels_like"`
	MinTemp       float32 `json:"temp_min"`
	MaxTemp       float32 `json:"temp_max"`
	Humidity      float32 `json:"humidity"`
	Pressure      float32 `json:"pressure"`
}

type wind struct {
	Speed   float32 `json:"speed"`
	Degrees float32 `json:"deg"`
	Gust    float32 `json:"gust"`
}

type description struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Weather struct {
	Core           weatherCore   `json:"main"`
	Wind           wind          `json:"wind"`
	Description    []description `json:"weather"`
	UnitSystemUsed units.UnitSystem
}

func (w *Weather) Describe() string {
	tempUnit := w.UnitSystemUsed.GetTempUnit()

	descriptions := ""
	for i, d := range w.Description {
		ending := ""
		if i != len(w.Description)-1 {
			ending = ","
		}
		descriptions += d.Description + ending
	}
	return fmt.Sprintf("It is %g %s (%g %s), %s", w.Core.Temp, tempUnit, w.Core.TempFeelsLike, tempUnit, descriptions)

}

func (w *Weather) DescribeDetails() string {
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()

	return fmt.Sprintf("Humidity: %g%s | Pressure: %g %s | Wind speed: %g %s", w.Core.Humidity, "%", w.Core.Pressure, pressureUnit, w.Wind.Speed, speedUnit)

}
