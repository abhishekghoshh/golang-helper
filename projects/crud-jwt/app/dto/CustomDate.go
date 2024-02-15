package dto

import (
	"encoding/json"
	"strings"
	"time"
)

type CustomDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *CustomDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = CustomDate(t)
	return nil
}
func (j CustomDate) MarshalJSON() ([]byte, error) {
	// fmt.Println(time.Time(j))
	return json.Marshal(time.Time(j))
}

func (j CustomDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
