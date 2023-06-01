# jsontime

A wrapper for the standard library [time.Time](https://pkg.go.dev/time) type that allows you to easily set formatting for marshaling and unmarshaling to JSON.

# Usage

## Converting to JSON

```go
type SampleTime struct {
	Time JsonTime `json:"time"`
}

sampleTime := SampleTime{Time: JsonTime(time.Now())}

jsonBody, err := json.Marshal(sampleTime)
fmt.Println(string(jsonBody)) // {"time":"2006-01-02T15:04:05.000Z"}

```

## Converting from JSON

```go
type SampleTime struct {
	Time JsonTime `json:"time"`
}

var sampleTime SampleTime

// JSON received from some API
jsonBody := []byte(`{"time": "2023-05-29T14:40:42.057Z"}`)
json.Unmarshal(jsonBody, &sampleTime)
```
