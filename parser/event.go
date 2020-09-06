package parser

import (
	"io"
	"encoding/binary"
)

type Event interface {
	GetDelta () uint64
}

type TrackEvent struct {
	delta uint64
}

type MetaEvent struct {
	TrackEvent
	// tmp
	data    []byte
	message byte
}

type SysExEvent struct {
	TrackEvent
	// tmp
	data    []byte
	message byte
}

type MidiEvent struct {
	TrackEvent
}

func (ev *TrackEvent) GetDelta () uint64 {
	return ev.delta
}

func ParseSysExEvent (delta uint64, evt byte, reader io.ByteReader) (*SysExEvent, error) {
	event := &SysExEvent {}
	event.TrackEvent.delta = delta

	messLen, err := binary.ReadUvarint(reader)
	if err != nil {
		panic(err)
	}

	event.data = make([]byte, messLen)

	for _, i := range event.data {
		data, err := reader.ReadByte()

		if err != nil {
			panic(err)
		}
		event.data[i] = data
	}

	return event, nil
}

func ParseMetaEvent (delta uint64, evt byte, reader io.ByteReader) (*MetaEvent, error) {
	event := &MetaEvent {}
	event.TrackEvent.delta = delta

	var err error

	event.message, err = reader.ReadByte()

	if err != nil {
		panic(err)
	}

	messLen, err := binary.ReadUvarint(reader)
	if err != nil {
		panic(err)
	}

	event.data = make([]byte, messLen)

	for _, i := range event.data {
		data, err := reader.ReadByte()

		if err != nil {
			panic(err)
		}
		event.data[i] = data
	}

	return event, nil
}

func ParseMidiEvent (delta uint64, evt byte, reader io.Reader) (*MidiEvent, error) {
	return nil, nil
}
