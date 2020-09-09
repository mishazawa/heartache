package events

type MidiEvent struct {
	DeltaTime
	status byte
	data   []byte
}

func ParseMidiEvent (evt *IntermediateEvent) *MidiEvent {
	return &MidiEvent{
		DeltaTime: DeltaTime(evt.delta),
		status: evt.status,
		data: evt.data,
	}
}
