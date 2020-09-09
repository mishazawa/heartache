package events

type SysExEvent struct {
	DeltaTime
	data    []byte
	message byte
}

func ParseSysExEvent (evt *IntermediateEvent) *SysExEvent {
	event := &SysExEvent {}
	return event
}
