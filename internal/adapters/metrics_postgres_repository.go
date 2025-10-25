package adapters

import (
	"context"

	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/not4sure/tracking-service/internal/common/db"
	"github.com/not4sure/tracking-service/internal/domain/metric"
)

type MetricsPostgresRepository struct {
	conn *pgxpool.Pool
}

func NewMetricsPostgresRepository(conn *pgxpool.Pool) MetricsPostgresRepository {
	return MetricsPostgresRepository{
		conn,
	}
}

func (r *MetricsPostgresRepository) Store(ctx context.Context, m *metric.Metric) error {
	arg := db.UpsertUserActivityMetricParams{
		UserID:        int64(m.UserID()),
		EventCount:    int32(m.EventCount()),
		WindowStartAt: convert.TimeToTimestamp(m.TimeWindow().Start()),
		WindowEndAt:   convert.TimeToTimestamp(m.TimeWindow().End()),
		CreatedAt:     convert.TimeToTimestamp(m.CreatedAt()),
	}

	return r.queries().UpsertUserActivityMetric(ctx, arg)
}

func (r *MetricsPostgresRepository) List(ctx context.Context) ([]*metric.Metric, error) {
	panic("uniplemented")
}

func (r *MetricsPostgresRepository) queries() *db.Queries {
	return db.New(r.conn)
}
