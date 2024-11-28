package weather

import (
	"fmt"
	"sort"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/pafello/gocast/internal/styles"
	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/utils"
)

type GroupedForecast map[time.Time]WeatherSeries

type Forecast struct {
	List           []Weather `json:"list"`
	UnitSystemUsed units.UnitSystem
}

func (f *Forecast) Describe(cityName string) string {
	desc := fmt.Sprintf("%s\n", cityName)

	for _, weather := range f.List {
		weather.UnitSystemUsed = f.UnitSystemUsed
		desc += weather.DescribeShort()
		desc += "\n"
	}
	return desc
}

func (f *Forecast) GroupWeatherByDay() GroupedForecast {

	groups := make(GroupedForecast)
	for _, w := range f.List {
		w.UnitSystemUsed = f.UnitSystemUsed

		dayMonth := utils.FormatTimeDayMonthYear(time.Unix(w.UnixTimestamp, 0))
		date, err := time.Parse(utils.DayMonthYearLayout, dayMonth)
		if err != nil {
			continue
		}

		groups[date] = append(groups[date], &w)

	}

	return groups
}

func (gf *GroupedForecast) DisplayDailyTable() {
	avgWeatherPerDay := make(map[time.Time]AvgDayWeather)

	for key, val := range *gf {
		avgWeather := val.GetAverageWeather()
		avgWeatherPerDay[key] = avgWeather
	}

	sortedKeys := []time.Time{}

	for key := range avgWeatherPerDay {
		sortedKeys = append(sortedKeys, key)
	}

	sort.Slice(sortedKeys, func(i, j int) bool {
		d1, d2 := sortedKeys[i], sortedKeys[j]
		return d1.Before(d2)
	})
	rows := [][]string{}

	for _, item := range sortedKeys {
		w := avgWeatherPerDay[item]

		tempUnit := w.UnitSystemUsed.GetTempUnit()
		pressureUnit := w.UnitSystemUsed.GetPressureUnit()
		speedUnit := w.UnitSystemUsed.GetSpeedUnit()
		tempLo := utils.LeftPad(fmt.Sprintf("%.2f %s", w.TempLo, tempUnit), 9)
		tempHi := utils.LeftPad(fmt.Sprintf("%.2f %s", w.TempHi, tempUnit), 9)
		pressure := utils.LeftPad(fmt.Sprintf("%.2f %s", w.Pressure, pressureUnit), 8)
		wind := utils.LeftPad(fmt.Sprintf("%.2f %s", w.WindSpeed, speedUnit), 8)
		rowData := []string{utils.FormatTimeDayMonth(item), tempLo, tempHi, pressure, wind}
		rows = append(rows, rowData)

	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(styles.TableBorder).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return styles.TableHeader
			case row%2 == 0:
				return styles.TableRowEven.Padding(0, 1, 0, 1)
			default:
				return styles.TableRowOdd.Padding(0, 1, 0, 1)
			}
		}).
		BorderBottom(true).
		BorderRow(true).
		Headers("Day", "Min temp", "Max temp", "Pressure", "Wind").
		Rows(rows...)

	fmt.Println(t)
}
