package data

import (
	"fmt"
	"time"
)

var (
	displayFormat = time.RFC1123
	loadFormat    = "2006-01-02T15:04:05-0700"
	localLocation = "America/Chicago"
)

// they server time looks to be -1 DT rn

type Event struct {
	ID    string `json:"idEvent"`
	Title string `json:"strEvent"`
	Date  string `json:"dateEvent"`
	Time  string `json:"strTime"`
}

func (e Event) IsToday() bool {
	loc, _ := time.LoadLocation(localLocation)
	if time.Now().In(loc).YearDay() == getTime(e.Date, e.Time).YearDay() {
		return true
	}

	return false
}

func getTime(date, ti string) time.Time {
	if len(date) > 0 && len(ti) > 0 {
		str := fmt.Sprint(date, "T", ti, "-0100")
		fmt.Println(str)
		t, _ := time.Parse(loadFormat, str)
		loc, _ := time.LoadLocation(localLocation)
		return t.In(loc)
	}

	return time.Time{}
}

func (e Event) GetTime() string {
	t := getTime(e.Date, e.Time)
	return t.Format(displayFormat)
}

type respEvents struct {
	Events []Event `json:"events"`
}

type Events []Event

func (e Events) Len() int      { return len(e) }
func (e Events) Swap(i, j int) { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool {
	ta := getTime(e[i].Date, e[i].Time)
	tb := getTime(e[j].Date, e[j].Time)
	if ta.Before(tb) {
		return true
	}

	return false
}
