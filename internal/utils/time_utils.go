package utils

import "time"

func FormatTime(t time.Time) string {

	return t.Format("_2 Jan 15:00")

}

const DayMonthYearLayout = "_2 Jan 2006"
const DayMonthLayout = "_2 Jan"

func FormatTimeDayMonthYear(t time.Time) string {
	return t.Format(DayMonthYearLayout)
}
func FormatTimeDayMonth(t time.Time) string {
	return t.Format(DayMonthLayout)
}
