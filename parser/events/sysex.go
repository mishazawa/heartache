package events

import (
	"io"
	"encoding/binary"
)

type SysExEvent struct {
	TrackEvent
	// tmp
	data    []byte
	message byte
}

func ParseSysExEvent (evt byte, reader io.ByteReader) (*SysExEvent, error) {
	event := &SysExEvent {}

	messLen, err := binary.ReadUvarint(reader)
	if err != nil {
		return event, err
	}

	event.data = make([]byte, messLen)

	for _, i := range event.data {
		data, err := reader.ReadByte()

		if err != nil {
			return event, err
		}
		event.data[i] = data
	}

	return event, nil
}
