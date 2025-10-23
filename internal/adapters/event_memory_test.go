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

func TestStoreEvent_OK(t *testing.T) {
	mr := adapters.NewEventMemoryRepository()
	e := randomEvent(t)

	err := mr.Store(context.Background(), e)

	require.NoError(t, err)
}

func TestStoreEvent_AlreadyExists(t *testing.T) {
	mr := adapters.NewEventMemoryRepository()
	e := randomEvent(t)
	err := mr.Store(context.Background(), e)
	require.NoError(t, err)

	err = mr.Store(context.Background(), e)

	require.Error(t, err)
	require.Equal(t, event.ErrEventAlreadyExists, err)
}

func TestFindByUUID_OK(t *testing.T) {
	mr := adapters.NewEventMemoryRepository()
	e := randomEvent(t)
	err := mr.Store(context.Background(), e)
	require.NoError(t, err)

	found, err := mr.FindByUUID(context.Background(), e.UUID())

	require.NoError(t, err)
	require.Equal(t, e, found)
}

func TestFindByUUID_NotFound(t *testing.T) {
	mr := adapters.NewEventMemoryRepository()

	e, err := mr.FindByUUID(context.Background(), uuid.New())

	require.Error(t, err)
	require.Equal(t, event.ErrEventNotFound, err)
	require.Nil(t, e)
}

func TestListEvents(t *testing.T) {
	mr := adapters.NewEventMemoryRepository()

	userID := randomUserID()
	const eventCount = 256
	for range eventCount {
		e := randomEventForUser(t, userID)
		err := mr.Store(context.Background(), e)
		require.NoError(t, err)
	}

	ee, err := mr.List(context.Background(), userID, time.Now().Add(-time.Minute), time.Now())

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
