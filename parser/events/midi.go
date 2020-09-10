package events

type MidiEvent struct {
	DeltaTime
	Status byte
	Data   []byte
}

func ParseMidiEvent (evt *IntermediateEvent) *MidiEvent {
	return &MidiEvent{
		DeltaTime: DeltaTime(evt.delta),
		Status: evt.status,
		Data: evt.data,
	}
}
