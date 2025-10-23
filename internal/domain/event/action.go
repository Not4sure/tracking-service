package event

type Action struct {
	a string
}

func (a Action) String() string {
	return a.a
}
