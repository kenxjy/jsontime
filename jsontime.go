package jsontime

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type JsonTime time.Time

// Default layout used
const JavaScriptToJSONFormat string = "2006-01-02T15:04:05.000Z"

var marshalLayout string = JavaScriptToJSONFormat
var unmarshalLayout string = JavaScriptToJSONFormat

/*
Override the layout used by MarshalJSON and UnmarshalJSON
See https://gosamples.dev/date-time-format-cheatsheet/ for writing layouts
*/
func SetLayout(layout string) {
	marshalLayout = layout
	unmarshalLayout = layout
}

/*
Override the layout used by MarshalJSON
See https://gosamples.dev/date-time-format-cheatsheet/ for writing layouts
*/
func SetMarshalLayout(layout string) {
	marshalLayout = layout
}

/*
Override the layout used by UnmarshalJSON
See https://gosamples.dev/date-time-format-cheatsheet/ for writing layouts
*/
func SetUnmarshalLayout(layout string) {
	unmarshalLayout = layout
}

func (jt *JsonTime) String() string {
	t := time.Time(*jt)
	return t.Format(marshalLayout)
}

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, jt.String())), nil
}

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	timeString := strings.Trim(string(b), `"`)
	t, err := time.Parse(unmarshalLayout, timeString)
	if err != nil {
		errorText := fmt.Sprintf("Invalid date format: %s", timeString)
		return errors.New(errorText)
	}

	*jt = JsonTime(t)
	return nil
}

func (jt *JsonTime) Time() time.Time {
	return time.Time(*jt)
}
