package helpers

import "time"

const Layout = "2006-01-02 15:04:05"

func StrToTime(date string) time.Time {
	t, _ := time.Parse(Layout, date)
	return t
}
