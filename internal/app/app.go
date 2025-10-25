package app

import (
	"github.com/not4sure/tracking-service/internal/app/command"
	"github.com/not4sure/tracking-service/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateEvent      command.CreateEventHandler
	CalculateMetrics command.CalculateMetricsHandler
}

type Queries struct {
	ListEvents  query.ListEventsHandler
	ListMetrics query.ListMetricsHandler
}
