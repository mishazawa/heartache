package events

type MidiEvent struct {
	DeltaTime
	Status byte
	Data   []byte
}

func ParseMidiEvent (evt *IntermediateEvent) *MidiEvent {
	event := &MidiEvent{}
	return event
}
