package utils

import "time"

func FormatTime(t time.Time) string {

	return t.Format("_2 Jan 15:00")

}

func FormatTimeDayMonth(t time.Time) string {
	return t.Format("_2 Jan")
}
