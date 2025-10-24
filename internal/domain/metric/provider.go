package metric

import (
	"context"
)

type Provider interface {
	AtTimeWindow(ctx context.Context, w TimeWindow) ([]*Metric, error)
}
