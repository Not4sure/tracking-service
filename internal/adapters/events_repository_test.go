package adapters_test

import (
	"context"
	"math/rand/v2"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/not4sure/tracking-service/internal/adapters"
	"github.com/not4sure/tracking-service/internal/domain/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	t.Parallel()

	repositories := createRepositories(t)

	for i := range repositories {
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("testStoreEvent_OK", func(t *testing.T) {
				t.Parallel()
				testStoreEvent_OK(t, r.Repository)
			})
			t.Run("testStoreEvent_AlreadyExists", func(t *testing.T) {
				t.Parallel()
				testStoreEvent_AlreadyExists(t, r.Repository)
			})
			t.Run("testFindByUUID_OK", func(t *testing.T) {
				t.Parallel()
				testFindByUUID_OK(t, r.Repository)
			})
			t.Run("testFindByUUID_NotFound", func(t *testing.T) {
				t.Parallel()
				testFindByUUID_NotFound(t, r.Repository)
			})
			t.Run("testListEvents", func(t *testing.T) {
				t.Parallel()
				testListEvents(t, r.Repository)
			})
		})
	}
}

type Repository struct {
	Name       string
	Repository event.Repository
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "Memory",
			Repository: adapters.NewEventsMemoryRepository(),
		},
		{
			Name:       "Postgres",
			Repository: newPostgresRepository(t, context.Background()),
		},
	}
}

func newPostgresRepository(t *testing.T, ctx context.Context) event.Repository {
	conn, err := adapters.NewPostgresConnection(ctx)
	require.NoError(t, err)

	return adapters.NewEventsPostgresRepository(conn)
}

func testStoreEvent_OK(t *testing.T, r event.Repository) {
	e := randomEvent(t)

	err := r.Store(context.Background(), e)

	require.NoError(t, err)
}

func testStoreEvent_AlreadyExists(t *testing.T, r event.Repository) {
	e := randomEvent(t)
	err := r.Store(context.Background(), e)
	require.NoError(t, err)

	err = r.Store(context.Background(), e)

	require.Error(t, err)
	require.Equal(t, event.ErrEventAlreadyExists, err)
}

func testFindByUUID_OK(t *testing.T, r event.Repository) {
	e := randomEvent(t)
	err := r.Store(context.Background(), e)
	require.NoError(t, err)

	found, err := r.FindByUUID(context.Background(), e.UUID())

	require.NoError(t, err)
	require.Equal(t, e.UUID(), found.UUID())
	require.WithinDuration(t, e.OccuredAt(), found.OccuredAt(), time.Second)
	require.Equal(t, e.UserID(), found.UserID())
	require.Equal(t, e.Action(), found.Action())
	require.Equal(t, e.Metadata(), found.Metadata())
}

func testFindByUUID_NotFound(t *testing.T, r event.Repository) {
	e, err := r.FindByUUID(context.Background(), uuid.New())

	require.Error(t, err)
	require.Equal(t, event.ErrEventNotFound, err)
	require.Nil(t, e)
}

func testListEvents(t *testing.T, r event.Repository) {
	userID := randomUserID()
	const eventCount = 256
	for range eventCount {
		e := randomEventForUser(t, userID)
		err := r.Store(context.Background(), e)
		require.NoError(t, err)
	}

	ee, err := r.List(context.Background(), userID, time.Now().Add(-time.Minute), time.Now())

	assert.Equal(t, eventCount, len(ee))
	require.NoError(t, err)
}

func randomEvent(t *testing.T) *event.Event {
	return randomEventForUser(t, randomUserID())
}

func randomEventForUser(t *testing.T, userID uint) *event.Event {
	return mustNewEvent(t, userID, randomAction(), randomMetadata())
}

func mustNewEvent(t *testing.T, userID uint, action string, metadata event.Metadata) *event.Event {
	e, err := event.New(userID, action, event.WithMetadata(metadata))
	require.NoError(t, err)

	return e
}

func randomUserID() uint {
	return rand.Uint()
}

func randomAction() string {
	return randomStringOf(64)
}

func randomMetadata() event.Metadata {
	m := event.Metadata{}
	for range rand.IntN(8) {
		m[randomStringOf(16)] = randomStringOf(64)
	}

	return m
}

func randomStringOf(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	return string(b)
}
