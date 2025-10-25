package metric

import (
	"context"
	"time"
)

type Repository interface {
	Store(ctx context.Context, m *Metric) error
	List(ctx context.Context, userID uint, from time.Time, till time.Time) ([]*Metric, error)
}
