package event

type EventOption func(*Event)

// WithMetadata sets given metadata for new Event.
func WithMetadata(metadata map[string]string) EventOption {
	return func(e *Event) {
		e.metadata = metadata
	}
}
