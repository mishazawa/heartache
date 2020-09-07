package events

import (
	"io"
)

type MidiEvent struct {
	TrackEvent
	status    byte
	key       byte
	velocity  byte
}

func ParseMidiEvent (evt byte, reader io.ByteReader) (*MidiEvent, error) {
	event := &MidiEvent{}
	event.status = evt

	var err error

	event.key, err = reader.ReadByte()
	if err != nil {
		return event, err
	}

	event.velocity, err = reader.ReadByte()
	if err != nil {
		return event, err
	}

	return event, nil
}
