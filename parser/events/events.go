package events

type Event interface {
	GetDelta () uint64
	SetDelta (uint64)
}

type TrackEvent struct {
	delta uint64
}

func (ev *TrackEvent) GetDelta () uint64 {
	return ev.delta
}

func (ev *TrackEvent) SetDelta (delta uint64) {
	ev.delta = delta
}
