package adapters

import (
	"context"

	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/not4sure/tracking-service/internal/common/db"
	"github.com/not4sure/tracking-service/internal/domain/metric"
)

type MetricsPostgresProvider struct {
	conn *pgxpool.Pool
}

func NewMetricsPostgresProvider(conn *pgxpool.Pool) MetricsPostgresProvider {
	return MetricsPostgresProvider{
		conn: conn,
	}
}

func (mp *MetricsPostgresProvider) AtTimeWindow(ctx context.Context, tw metric.TimeWindow) ([]*metric.Metric, error) {
	arg := db.CountEventsByUserParams{
		OccuredAt:   convert.TimeToTimestamp(tw.Start()),
		OccuredAt_2: convert.TimeToTimestamp(tw.End()),
	}

	rows, err := db.New(mp.conn).CountEventsByUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	mm := []*metric.Metric{}
	for _, r := range rows {
		mm = append(mm, metric.NewMetric(uint(r.UserID), uint(r.EventCount), tw))
	}

	return mm, nil
}
