package api

import (
	"encoding/json"
	"time"
)

func parseDate(d string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", d, time.UTC)
	if err != nil {
		panic(err)
	}
	return t
}

func toJSON(payload interface{}) []byte {
	js, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return js
}
