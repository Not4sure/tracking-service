package metric

import (
	"fmt"
	"time"
)

var timeWindowDuration = 4 * time.Hour

type TimeWindow struct {
	start time.Time
}

func (t TimeWindow) Start() time.Time {
	return t.start
}

func (t TimeWindow) End() time.Time {
	return t.start.Add(timeWindowDuration)
}

func (t TimeWindow) Prev() TimeWindow {
	return TimeWindow{
		start: t.start.Add(-timeWindowDuration),
	}
}

func PrevTimeWindow() TimeWindow {
	return CurrentTimeWindow().Prev()
}

func CurrentTimeWindow() TimeWindow {
	return TimeWindowAt(time.Now())
}

func TimeWindowAt(t time.Time) TimeWindow {
	return TimeWindow{
		start: getStartFor(t),
	}
}

func getStartFor(t time.Time) time.Time {
	t = t.UTC()
	startHour := (t.Hour() / 4) * 4

	return time.Date(t.Year(), t.Month(), t.Day(), startHour, 0, 0, 0, time.UTC)
}

func (t TimeWindow) String() string {
	return fmt.Sprintf(
		"TimeWindow Start: %s End: %s",
		t.Start().Format(time.RFC3339),
		t.End().Format(time.RFC3339),
	)
}
