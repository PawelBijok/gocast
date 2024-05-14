package weather

import (
	"fmt"
	"sort"
	"time"

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

func (gf *GroupedForecast) DescribeDaily() {
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

	for _, item := range sortedKeys {
		weather := avgWeatherPerDay[item]

		fmt.Printf("%s: %s\n", utils.FormatTimeDayMonth(item), weather.DescribeShort())
	}

}
