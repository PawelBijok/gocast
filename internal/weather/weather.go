package weather

import (
	"fmt"

	"github.com/pafello/gocast/internal/styles"
	"github.com/pafello/gocast/internal/units"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
type AvgDayWeather struct {
	TempLo         float32
	TempHi         float32
	Pressure       float32
	WindSpeed      float32
	UnitSystemUsed units.UnitSystem
}

func (w *AvgDayWeather) DescribeShort() string {
	tempUnit := w.UnitSystemUsed.GetTempUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	tempLo := utils.LeftPad(fmt.Sprintf("%.2f %s", w.TempLo, tempUnit), 9)
	tempHi := utils.LeftPad(fmt.Sprintf("%.2f %s", w.TempHi, tempUnit), 9)
	pressure := utils.LeftPad(fmt.Sprintf("%.2f %s", w.Pressure, pressureUnit), 8)
	wind := utils.LeftPad(fmt.Sprintf("%.2f %s", w.WindSpeed, speedUnit), 8)

	return fmt.Sprintf(" %s - %s  | %s | %s", tempLo, tempHi, pressure, wind)

}

type WeatherSeries []*Weather

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
	return fmt.Sprintf("%g %s (%g %s), %s", w.Core.Temp, tempUnit, w.Core.TempFeelsLike, tempUnit, descriptions)

}

func (w *Weather) DescribeShort() string {

	tempUnit := w.UnitSystemUsed.GetTempUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	temp := utils.LeftPad(fmt.Sprintf("%.2f %s", w.Core.Temp, tempUnit), 10)
	pressure := utils.LeftPad(fmt.Sprintf("%.2f %s", w.Core.Pressure, pressureUnit), 8)
	wind := utils.LeftPad(fmt.Sprintf("%.2f %s", w.Wind.Speed, speedUnit), 8)

	return fmt.Sprintf(" %s | %s | %s", temp, pressure, wind)

}

func (w *Weather) DisplayDetails() {
	descriptions := ""
	for i, d := range w.Description {
		ending := ""
		if i != len(w.Description)-1 {
			ending = ","
		}
		descriptions += d.Description + ending
	}
	tempUnit := w.UnitSystemUsed.GetTempUnit()
	speedUnit := w.UnitSystemUsed.GetSpeedUnit()
	pressureUnit := w.UnitSystemUsed.GetPressureUnit()

	rows := [][]string{

		{"Feels like", fmt.Sprintf("%.2f %s", w.Core.TempFeelsLike, tempUnit)},
		{"State", descriptions},
		{"Humidity", fmt.Sprintf("%g%s", w.Core.Humidity, "%")},
		{"Pressure", fmt.Sprintf("%g %s", w.Core.Pressure, pressureUnit)},
		{"Wind speed", fmt.Sprintf("%g %s", w.Wind.Speed, speedUnit)},
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(styles.TableBorder).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:

				return styles.TableHeader
			case row%2 == 0:
				if col == 0 {
					return styles.TableRowEven.Padding(0, 1, 0, 1).Align(lipgloss.Right)
				}
				return styles.TableRowEven.Padding(0, 1, 0, 1).Align(lipgloss.Left)
			default:
				if col == 0 {
					return styles.TableRowOdd.Padding(0, 1, 0, 1).Align(lipgloss.Right)
				}
				return styles.TableRowOdd.Padding(0, 1, 0, 1).Align(lipgloss.Left)
			}
		}).
		Headers("Temp", fmt.Sprintf("%.2f %s", w.Core.Temp, tempUnit)).
		BorderBottom(true).
		BorderRow(false).
		Rows(rows...)

	fmt.Println(t)

}

func (ws WeatherSeries) GetAverageWeather() AvgDayWeather {
	seriesQuantity := len(ws)
	var avgTemp float32 = 0
	var avgTempNight float32 = 0
	var avgPressure float32 = 0
	var avgWindSpeed float32 = 0
	var tempLo float32 = 1000
	var tempHi float32 = -1000

	for i := 0; i < seriesQuantity; i++ {

		w := ws[i]
		if w.Core.Temp > tempHi {
			tempHi = w.Core.Temp
		}
		if w.Core.Temp < tempLo {
			tempLo = w.Core.Temp
		}
		avgPressure += w.Core.Pressure
		avgWindSpeed += w.Wind.Speed
	}
	avgTemp /= float32(seriesQuantity)
	avgTempNight /= float32(seriesQuantity)
	avgPressure /= float32(seriesQuantity)
	avgWindSpeed /= float32(seriesQuantity)

	avgWeather := AvgDayWeather{Pressure: avgPressure, WindSpeed: avgWindSpeed, UnitSystemUsed: ws[0].UnitSystemUsed, TempLo: tempLo, TempHi: tempHi}

	return avgWeather
}
