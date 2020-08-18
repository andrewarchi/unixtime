// Package unixtime is a time.Time wrapper for marshalling Unix
// timestamps in JSON.
//
package unixtime

import (
	"strconv"
	"time"
)

// Time is a time formatted as a Unix timestamp.
type Time struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
// The time is a number representing a Unix timestamp.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a number representing a Unix timestamp.
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	sec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = Time{time.Unix(sec, 0)}
	return nil
}
