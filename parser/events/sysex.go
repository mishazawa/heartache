package events

type SysExEvent struct {
	DeltaTime
	data    []byte
	status byte
}

func ParseSysExEvent (evt *IntermediateEvent) *SysExEvent {
	return &SysExEvent {
		DeltaTime: DeltaTime(evt.delta),
		status: evt.status,
		data: evt.data,
	}
}
