package jsontime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type SampleTime struct {
	Time JsonTime `json:"time"`
}

func TestToTime(t *testing.T) {
	asTime, _ := time.Parse(JavaScriptToJSONFormat, "2002-05-29T14:40:42.057Z")
	jsonTime := JsonTime(asTime)

	got := fmt.Sprintf("%T", jsonTime.Time())
	want := "time.Time"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestUnmarshal(t *testing.T) {
	var sampleTime SampleTime

	jsonBody := []byte(`{"time": "2023-05-29T14:40:42.057Z"}`)
	json.Unmarshal(jsonBody, &sampleTime)

	got := sampleTime.Time.Time().UnixMilli()
	want := int64(1685371242057)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestMarshal(t *testing.T) {
	asTime, _ := time.Parse(JavaScriptToJSONFormat, "2002-05-29T14:40:42.057Z")
	sampleTime := SampleTime{JsonTime(asTime)}

	jsonBody, _ := json.Marshal(sampleTime)
	got := string(jsonBody)
	want := `{"time":"2002-05-29T14:40:42.057Z"}`

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSetLayout(t *testing.T) {
	SetLayout("2006-01-02")

	var sampleTime SampleTime

	jsonBody := []byte(`{"time": "2023-05-29"}`)
	json.Unmarshal(jsonBody, &sampleTime)

	gotYear := sampleTime.Time.Time().Year()
	wantYear := 2023

	if gotYear != wantYear {
		t.Errorf("got %d, wanted %d", gotYear, wantYear)
	}

	gotMonth := sampleTime.Time.Time().Month()
	wantMonth := time.May

	if gotMonth != wantMonth {
		t.Errorf("got %s, wanted %s", gotMonth, wantMonth)
	}

	gotDay := sampleTime.Time.Time().Day()
	wantDay := 29

	if gotDay != wantDay {
		t.Errorf("got %d, wanted %d", gotDay, wantDay)
	}
}

func TestSetMarshalLayout(t *testing.T) {
	SetLayout(JavaScriptToJSONFormat)
	SetMarshalLayout("2006-01-02")

	var sampleTime SampleTime

	originalJsonBody := []byte(`{"time": "2023-05-29T14:40:42.057Z"}`)
	json.Unmarshal(originalJsonBody, &sampleTime)
	newJsonBody, _ := json.Marshal(sampleTime)
	got := string(newJsonBody)
	want := `{"time":"2023-05-29"}`

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestSetUnmarshalLayout(t *testing.T) {
	SetLayout(JavaScriptToJSONFormat)
	SetUnmarshalLayout("2006-01-02")

	var sampleTime SampleTime

	originalJsonBody := []byte(`{"time": "2023-05-29"}`)
	json.Unmarshal(originalJsonBody, &sampleTime)
	newJsonBody, _ := json.Marshal(sampleTime)
	got := string(newJsonBody)
	want := `{"time":"2023-05-29T00:00:00.000Z"}`

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestLayoutError(t *testing.T) {
	SetLayout("2006-01-02")

	var sampleTime SampleTime

	jsonBody := []byte(`{"time": "2023-05-29T14:40:42.057Z"}`)
	err := json.Unmarshal(jsonBody, &sampleTime)

	gotErr := err == nil
	wantErr := true

	if gotErr {
		t.Errorf("err == nil: got %t, wanted %t", wantErr, gotErr)
	}

	got := err.Error()
	want := "Invalid date format: 2023-05-29T14:40:42.057Z"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
