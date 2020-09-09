package events

import (
	"encoding/binary"
)


type VariableMetaEvent struct {
	DeltaTime
	data   []byte
	status byte
}

type KeySignatureEvent struct {
	DeltaTime
	shift int8
	key   byte
}

type TimeSignatureEvent struct {
	DeltaTime
	nominator   uint8
	denominator uint8 // 2^denominetor
	clocks      uint8 // MIDI clocks
	beats       uint8 // 32nd notes per MIDI 1/4 note
}

type SMPTEOffsetEvent struct {
	DeltaTime
	hours     uint8
	minutes   uint8
	seconds   uint8
	frames    uint8
	subframes uint8
}

type SetTempoEvent struct {
	DeltaTime
	tempo uint32
}

type EndOfTrackEvent struct {
	DeltaTime
}

type MidiChannelPrefixEvent struct {
	DeltaTime
	channel byte
}

type MetaSequenceNumberEvent struct {
	DeltaTime
	seqnum1 byte
	seqnum2 byte
}

func ParseMetaEvent (evt *IntermediateEvent) Event {
	switch evt.status {
	case 0x00:
		return &MetaSequenceNumberEvent {
			DeltaTime: DeltaTime(evt.delta),
			seqnum1: evt.data[0],
			seqnum2: evt.data[1],
		}
	case 0x20:
		return &MidiChannelPrefixEvent {
			DeltaTime: DeltaTime(evt.delta),
			channel: evt.data[0],
		}
	case 0x2f:
		return &EndOfTrackEvent{ DeltaTime(evt.delta) }
	case 0x51:
		return &SetTempoEvent {
			DeltaTime: DeltaTime(evt.delta),
			tempo: padUint32(evt.data),
		}
	case 0x54:
		return &SMPTEOffsetEvent{
			DeltaTime: DeltaTime(evt.delta),
			hours:     evt.data[0],
			minutes:   evt.data[1],
			seconds:   evt.data[2],
			frames:    evt.data[3],
			subframes: evt.data[4],
		}
	case 0x58:
		return &TimeSignatureEvent{
			DeltaTime: DeltaTime(evt.delta),
			nominator:   evt.data[0],
			denominator: evt.data[1],
			clocks:      evt.data[2],
			beats:       evt.data[3],
		}
	case 0x59:
		return &KeySignatureEvent{
			DeltaTime: DeltaTime(evt.delta),
			shift: int8(evt.data[0]),
			key:   evt.data[1],
		}
	default:
		return &VariableMetaEvent{
			DeltaTime: DeltaTime(evt.delta),
			status: evt.status,
			data: evt.data,
		}
	}
	return nil
}

func padUint32(s []byte) uint32 {
	var b [4]byte
	copy(b[4-len(s):], s)
	return binary.BigEndian.Uint32(b[:])
}
