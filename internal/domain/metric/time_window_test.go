package metric_test

import (
	"testing"
	"time"

	"github.com/not4sure/tracking-service/internal/domain/metric"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeWindowAt(t *testing.T) {
	testCases := []struct {
		t     time.Time
		start time.Time
		end   time.Time
	}{
		{
			t:     mustParseKitchen(t, "2:02AM"),
			start: mustParseKitchen(t, "0:00AM"),
			end:   mustParseKitchen(t, "4:00AM"),
		},
		{
			t:     mustParseKitchen(t, "4:01AM"),
			start: mustParseKitchen(t, "4:00AM"),
			end:   mustParseKitchen(t, "8:00AM"),
		},
		{
			t:     mustParseKitchen(t, "11:59AM"),
			start: mustParseKitchen(t, "8:00AM"),
			end:   mustParseKitchen(t, "12:00PM"),
		},
		{
			t:     mustParseKitchen(t, "12:00PM"),
			start: mustParseKitchen(t, "12:00PM"),
			end:   mustParseKitchen(t, "4:00PM"),
		},
		{
			t:     mustParseKitchen(t, "5:00PM"),
			start: mustParseKitchen(t, "4:00PM"),
			end:   mustParseKitchen(t, "8:00PM"),
		},
		{
			t:     mustParseKitchen(t, "10:59PM"),
			start: mustParseKitchen(t, "8:00PM"),
			end:   mustParseKitchen(t, "12:00AM").Add(24 * time.Hour),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.t.Format(time.Kitchen), func(t *testing.T) {
			w := metric.TimeWindowAt(tc.t)

			assert.Equal(t, tc.start, w.Start())
			assert.Equal(t, tc.end, w.End())
		})
	}
}

func mustParseKitchen(t *testing.T, v string) time.Time {
	timestamp, err := time.Parse(time.Kitchen, v)
	require.NoError(t, err)

	return timestamp
}
