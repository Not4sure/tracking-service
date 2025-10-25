package metric

import "time"

type Metric struct {
	userID     uint
	eventCount uint
	timeWindow TimeWindow
	createdAt  time.Time
}

func NewMetric(userID uint, eventCount uint, w TimeWindow) *Metric {
	return &Metric{
		userID:     userID,
		eventCount: eventCount,
		timeWindow: w,
		createdAt:  time.Now(),
	}
}

// UnmarshalMetricFromDatabase unmarshals Metric the database.
//
// It should be used only for unmarshalling from the database!
// You can't use UnmarshalMetricFromDatabase as constructor - It may put domain into the invalid state!
func UnmarshalMetricFromDatabase(
	userID uint,
	eventCount uint,
	timeWindowStart time.Time,
	createdAt time.Time,
) *Metric {
	m := NewMetric(userID, eventCount, TimeWindow{start: timeWindowStart})

	m.createdAt = createdAt
	return m
}

func (m Metric) UserID() uint {
	return m.userID
}

func (m Metric) EventCount() uint {
	return m.eventCount
}

func (m Metric) TimeWindow() TimeWindow {
	return m.timeWindow
}

func (m Metric) CreatedAt() time.Time {
	return m.createdAt
}
