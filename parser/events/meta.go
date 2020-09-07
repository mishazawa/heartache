package events

import (
	"io"
	"encoding/binary"
)


type VariableMetaEvent struct {
	TrackEvent
	data    []byte
	message byte
}

type KeySignatureEvent struct {
	TrackEvent
	shift int8
	key   byte
}

type TimeSignatureEvent struct {
	TrackEvent
	nominator   uint8
	denominator uint8 // 2^denominetor
	clocks      uint8 // MIDI clocks
	beats       uint8 // 32nd notes per MIDI 1/4 note
}

type SMPTEOffsetEvent struct {
	TrackEvent
	hours     uint8
	minutes   uint8
	seconds   uint8
	frames    uint8
	subframes uint8
}

type SetTempoEvent struct {
	TrackEvent
	tempo uint32
}

type EndOfTrackEvent struct {
	TrackEvent
}

type MidiChannelPrefixEvent struct {
	TrackEvent
	channel byte
}

type MetaSequenceNumberEvent struct {
	TrackEvent
	seqnum1 byte
	seqnum2 byte
}

func ParseMetaEvent (evt byte, reader io.ByteReader) (Event, error) {
	message, err := reader.ReadByte()

	if err != nil {
		return nil, err
	}

	var event Event

	switch message {
	case 0x00:
		event, err = parseSequenceNumberEvent(reader)
	case 0x20:
		event, err = parseMidiChannelPrefixEvent(reader)
	case 0x2f:
		event, err = parseEndOfTrackEvent(reader)
	case 0x51:
		event, err = parseSetTempoEvent(reader)
	case 0x54:
		event, err = parseSMPTEOffsetEvent(reader)
	case 0x58:
		event, err = parseTimeSignatureEvent(reader)
	case 0x59:
		event, err = parseKeySignatureEvent(reader)
	default:
		event, err = parseVariableMetaEvent(message, reader)
	}

	return event, nil
}

func parseVariableMetaEvent (message byte, reader io.ByteReader) (*VariableMetaEvent, error) {
	event := &VariableMetaEvent{}

	messageLen, err := binary.ReadUvarint(reader)

	if err != nil {
		return nil, err
	}

	event.data = make([]byte, messageLen)

	for _, i := range event.data {
		data, err := reader.ReadByte()

		if err != nil {
			return nil, err
		}

		event.data[i] = data
	}

	event.message = message
	return event, nil
}

func parseKeySignatureEvent (reader io.ByteReader) (*KeySignatureEvent, error) {
	var err error

	_, err = reader.ReadByte() // skip len
	if err != nil {
		return nil, err
	}

	var shiftData byte
	shiftData, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}


	var key byte
	key, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}

	event := &KeySignatureEvent{
		shift: int8(shiftData),
		key:   key,
	}

	return event, err
}

func parseTimeSignatureEvent (reader io.ByteReader) (*TimeSignatureEvent, error) {
	tsLen, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, tsLen)

	for i := range buf {
		data, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		buf[i] = data
	}
	event := &TimeSignatureEvent{
		nominator:   buf[0],
		denominator: buf[1],
		clocks:      buf[2],
		beats:       buf[3],
	}

	return event, nil
}

func parseSMPTEOffsetEvent (reader io.ByteReader) (*SMPTEOffsetEvent, error) {

	offsetLen, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, offsetLen)

	for i := range buf {
		data, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		buf[i] = data
	}

	event := &SMPTEOffsetEvent{
		hours:     buf[0],
		minutes:   buf[1],
		seconds:   buf[2],
		frames:    buf[3],
		subframes: buf[4],
	}

	return event, nil
}

func parseSetTempoEvent (reader io.ByteReader) (*SetTempoEvent, error) {
	event := &SetTempoEvent {}

	_, err := reader.ReadByte() // skip length
	if err != nil {
		return nil, err
	}

	// 32 bit just to convert byte array to uint32
	buf := []byte {0x00, 0x00, 0x00, 0x00}

	for i := 1; i < len(buf); i += 1 {
		data, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		buf[i] = data
	}

	event.tempo = binary.BigEndian.Uint32(buf)
	return event, err
}

func parseEndOfTrackEvent (reader io.ByteReader) (*EndOfTrackEvent, error) {
	event := &EndOfTrackEvent {}
	_, err := reader.ReadByte() // no other bytes in message
	return event, err
}

func parseMidiChannelPrefixEvent (reader io.ByteReader) (*MidiChannelPrefixEvent, error) {
	var err error

	event := &MidiChannelPrefixEvent {}

	_, err = reader.ReadByte() // skip length

	if err != nil {
		return nil, err
	}

	event.channel, err = reader.ReadByte()

	return event, err
}

func parseSequenceNumberEvent (reader io.ByteReader) (*MetaSequenceNumberEvent, error) {
	_, err := reader.ReadByte() // skip length
	if err != nil {
		return nil, err
	}

	event := &MetaSequenceNumberEvent {}
	event.seqnum1, err = reader.ReadByte()

	if err != nil {
		return nil, err
	}

	event.seqnum2, err = reader.ReadByte()
	if err != nil {
		return nil, err
	}

	return event, nil
}
