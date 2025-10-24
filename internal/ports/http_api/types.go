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
