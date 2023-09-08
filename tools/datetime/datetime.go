package datetime

import (
	"fmt"
	"time"
)

const (
	DefaultDateFormat     = "2006-01-02"
	DefaultDateTimeFormat = "2006-01-02 15:04:05.000Z"
)

var (
	ErrInvalidDateFormat     = fmt.Errorf("invalid date format (%s)", DefaultDateFormat)
	ErrInvalidDateTimeFormat = fmt.Errorf("invalid datetime format (%s)", DefaultDateTimeFormat)
)

func ParseDate(s string) (time.Time, error) {
	t, err := time.Parse(DefaultDateFormat, s)
	if err != nil {
		return t, ErrInvalidDateFormat
	}

	return t, nil
}
