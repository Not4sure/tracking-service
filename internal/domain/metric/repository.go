package metric

import "context"

type Repository interface {
	Store(ctx context.Context, m *Metric) error
}
