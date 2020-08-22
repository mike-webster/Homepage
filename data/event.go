package data

import (
	"fmt"
	"time"
)

// they server time looks to be -1 DT rn

type Event struct {
	ID    string `json:"idEvent"`
	Title string `json:"strEvent"`
	Date  string `json:"dateEvent"`
	Time  string `json:"strTime"`
}

func (e Event) GetTime() string {
	format := "2006-01-02T15:04:05-0700"
	if len(e.Date) > 0 && len(e.Time) > 0 {
		str := fmt.Sprint(e.Date, "T", e.Time, "-0100")
		fmt.Println(str)
		t, _ := time.Parse(format, str)
		loc, _ := time.LoadLocation("America/Chicago")
		return t.In(loc).Format(format)
	}

	return time.Time{}.Format(format)
}

type respEvents struct {
	Events []Event `json:"events"`
}

type Events []Event

func (e Events) Len() int      { return len(e) }
func (e Events) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool {
	if e[i].Time <= e[j].Time {
		return true
	}

	return false
}
