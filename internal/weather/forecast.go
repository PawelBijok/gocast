package weather

import (
	"fmt"
	"time"

	"github.com/pafello/gocast/internal/units"
	"github.com/pafello/gocast/internal/utils"
)

type GroupedForecast map[string]WeatherSeries

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
		dayMonth := utils.FormatTimeDayMonth(time.Unix(w.UnixTimestamp, 0))

		groups[dayMonth] = append(groups[dayMonth], &w)

	}

	return groups
}

func (gf *GroupedForecast) DescribeDaily() {
	avgWeatherPerDay := make(map[string]Weather)

	for key, val := range *gf {
		avgWeather := val.GetAverageWeather()
		avgWeatherPerDay[key] = avgWeather
		fmt.Println(avgWeather.DescribeShort())
	}
}
