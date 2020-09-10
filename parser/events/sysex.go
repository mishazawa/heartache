package events

type SysExEvent struct {
	DeltaTime
	Data    []byte
	Status byte
}

func ParseSysExEvent (evt *IntermediateEvent) *SysExEvent {
	return &SysExEvent {
		DeltaTime: DeltaTime(evt.delta),
		Status: evt.status,
		Data: evt.data,
	}
}
