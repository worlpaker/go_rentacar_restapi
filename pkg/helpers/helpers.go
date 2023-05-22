package helpers

import (
	Log "backend/pkg/helpers/log"
	"time"
)

// DateToInt converts a given slice of date strings (format "2006-01-02")
func DateToInt(date ...string) (tInt []int, err error) {
	dateFormat := "2006-01-02"
	for i := range date {
		t, err := time.Parse(dateFormat, date[i])
		if Log.Err(err) {
			return nil, err
		}
		tInt = append(tInt, int(t.Weekday()))
	}
	return
}

// IsParamNull checks whether any of the given parameters is an empty string.
func IsParamNull(param ...string) (ok bool) {
	for i := range param {
		if param[i] == "" {
			ok = true
			return
		}
	}
	return
}
