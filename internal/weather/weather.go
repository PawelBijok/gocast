package weather

import (
	"fmt"
	"time"

	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/utils"
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
	UnixTimestamp  int64         `json:"dt"`
	UnitSystemUsed units.UnitSystem
}

type WeatherSeries []*Weather

func (w *Weather) Describe(cityName string) string {
	tempUnit := w.UnitSystemUsed.GetTempUnit()

	descriptions := ""
	for i, d := range w.Description {
		ending := ""
		if i != len(w.Description)-1 {
			ending = ","
		}
		descriptions += d.Description + ending
	}
	return fmt.Sprintf("%s: %g %s (%g %s), %s", cityName, w.Core.Temp, tempUnit, w.Core.TempFeelsLike, tempUnit, descriptions)

}

func (w *Weather) DescribeShort() string {
	tempUnit := w.UnitSystemUsed.GetTempUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	temp := utils.LeftPad(fmt.Sprintf("%g %s", w.Core.Temp, tempUnit), 10)
	pressure := utils.LeftPad(fmt.Sprintf("%g %s", w.Core.Pressure, pressureUnit), 8)
	wind := utils.LeftPad(fmt.Sprintf("%g %s", w.Wind.Speed, speedUnit), 8)

	time := utils.FormatTime(time.Unix(w.UnixTimestamp, 0))
	return fmt.Sprintf("%s: %s | %s | %s", time, temp, pressure, wind)

}

func (w *Weather) DescribeDetails() string {
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()

	return fmt.Sprintf("Humidity: %g%s | Pressure: %g %s | Wind speed: %g %s", w.Core.Humidity, "%", w.Core.Pressure, pressureUnit, w.Wind.Speed, speedUnit)

}

func (ws WeatherSeries) GetAverageWeather() Weather {
	seriesQuantity := len(ws)
	var avgTemp float32 = 0
	for i := 0; i < seriesQuantity; i++ {

		w := ws[i]
		avgTemp += (w.Core.Temp)
	}
	avgTemp /= float32(seriesQuantity)
	avgWeather := ws[0]
	avgWeather.Core.Temp = avgTemp
	return *avgWeather
}
