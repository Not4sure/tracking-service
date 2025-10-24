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
