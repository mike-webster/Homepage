package data

import "time"

type Event struct {
	ID    string    `json:"idEvent"`
	Title string    `json:"strEvent"`
	Time  time.Time `json:"strTimestamp" time_format:"2006-01-02T15:04:05-07:00"`
}

type respEvents struct {
	Events []Event `json:"events"`
}
