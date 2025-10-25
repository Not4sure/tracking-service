package http_api

import "time"

type Events struct {
	Events []Event `json:"events"`
}

type Event struct {
	UUID      string    `json:"uuid"`
	OccuredAt time.Time `json:"occured_at"`

	UserID   uint              `json:"user_id"`
	Action   string            `json:"action"`
	Metadata map[string]string `json:"metadata"`
}

type Metrics struct {
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	UserID        uint      `json:"user_id"`
	EventCount    uint      `json:"event_count"`
	WindowStartAt time.Time `json:"window_start_at"`
	WindowEndAt   time.Time `json:"window_end_at"`
	CreatedAt     time.Time `json:"created_at"`
}
