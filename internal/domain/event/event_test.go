package event_test

import (
	"testing"
	"time"

	"github.com/not4sure/tracking-service/internal/domain/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewEvent_OK(t *testing.T) {
	testCases := []struct {
		name     string
		userID   uint
		action   string
		metadata map[string]string
	}{
		{
			name:   "view home page",
			userID: 1,
			action: "page_view",
			metadata: map[string]string{
				"page": "/home",
			},
		},
		{
			name:   "add item to cart",
			userID: 2,
			action: "add_to_cart",
			metadata: map[string]string{
				"item_id": "ea1f7f61-a2b6-4f3d-8a8a-0ae59d262786",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e, err := event.New(tc.userID, tc.action, event.WithMetadata(tc.metadata))

			assert.Equal(t, tc.userID, e.UserID())
			assert.Equal(t, tc.action, e.Action())
			assert.Equal(t, event.Metadata(tc.metadata), e.Metadata())
			assert.WithinDuration(t, time.Now(), e.OccuredAt(), time.Second)

			require.NoError(t, err)
		})
	}
}
