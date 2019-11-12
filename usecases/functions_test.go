package usecases_test

import "time"

func parseDate(d string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", d, time.UTC)
	if err != nil {
		panic(err)
	}
	return t
}
