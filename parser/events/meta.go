package events

import (
	"encoding/binary"
)

type VariableMetaEvent struct {
	DeltaTime
	Data   []byte
	Status byte
}

type KeySignatureEvent struct {
	DeltaTime
	Shift int8
	Key   byte
}

type TimeSignatureEvent struct {
	DeltaTime
	Nominator   uint8
	Denominator uint8 // 2^denominetor
	Clocks      uint8 // MIDI clocks
	Beats       uint8 // 32nd notes per MIDI 1/4 note
}

type SMPTEOffsetEvent struct {
	DeltaTime
	Hours     uint8
	Minutes   uint8
	Seconds   uint8
	Frames    uint8
	Subframes uint8
}

type SetTempoEvent struct {
	DeltaTime
	Tempo uint32
}

type EndOfTrackEvent struct {
	DeltaTime
}

type MidiChannelPrefixEvent struct {
	DeltaTime
	Channel byte
}

type MetaSequenceNumberEvent struct {
	DeltaTime
	SequenceNumber1 byte
	SequenceNumber2 byte
}

func ParseMetaEvent (evt *IntermediateEvent) Event {
	switch evt.status {
	case 0x00:
		return &MetaSequenceNumberEvent {
			DeltaTime: DeltaTime(evt.delta),
			SequenceNumber1: evt.data[0],
			SequenceNumber2: evt.data[1],
		}
	case 0x20:
		return &MidiChannelPrefixEvent {
			DeltaTime: DeltaTime(evt.delta),
			Channel: evt.data[0],
		}
	case 0x2f:
		return &EndOfTrackEvent{ DeltaTime(evt.delta) }
	case 0x51:
		return &SetTempoEvent {
			DeltaTime: DeltaTime(evt.delta),
			Tempo: padUint32(evt.data),
		}
	case 0x54:
		return &SMPTEOffsetEvent{
			DeltaTime: DeltaTime(evt.delta),
			Hours:     evt.data[0],
			Minutes:   evt.data[1],
			Seconds:   evt.data[2],
			Frames:    evt.data[3],
			Subframes: evt.data[4],
		}
	case 0x58:
		return &TimeSignatureEvent{
			DeltaTime: DeltaTime(evt.delta),
			Nominator:   evt.data[0],
			Denominator: evt.data[1],
			Clocks:      evt.data[2],
			Beats:       evt.data[3],
		}
	case 0x59:
		return &KeySignatureEvent{
			DeltaTime: DeltaTime(evt.delta),
			Shift: int8(evt.data[0]),
			Key:   evt.data[1],
		}
	default:
		return &VariableMetaEvent{
			DeltaTime: DeltaTime(evt.delta),
			Status: evt.status,
			Data: evt.data,
		}
	}
	return nil
}

func padUint32(s []byte) uint32 {
	var b [4]byte
	copy(b[4-len(s):], s)
	return binary.BigEndian.Uint32(b[:])
}
