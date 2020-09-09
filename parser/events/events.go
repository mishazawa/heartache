package events

import (
	"io"
	"fmt"
	"bytes"
	"encoding/binary"
)

const MIDI_MESSAGE  = "midi"
const META_MESSAGE  = "meta"
const SYSEX_MESSAGE = "sysex"

type DeltaTime uint64

func (dt *DeltaTime) GetDelta () *DeltaTime {
	return dt
}

type Event interface {
	GetDelta () *DeltaTime
}

type IntermediateEvent struct {
	message string
	delta   uint64
	status  byte
	data    []byte
}

func NewIntermediateEvent () *IntermediateEvent {
	return &IntermediateEvent{}
}

func (event *IntermediateEvent) ParseEvent (status byte, data []byte, delta uint64, r *bytes.Reader) error {
	event.data = data
	event.delta = delta
	event.status = status

	// check running status chunk
	if len(data) != 0 {
		switch status & 0xF0 {
		case 0xC0, 0xD0:
		default:
			nextByte, err := r.ReadByte()
			if err != nil {
				panic(err)
			}
			event.data = append(event.data, nextByte)
		}
		event.message = MIDI_MESSAGE
	} else {
		// regular messages
		if status & 0xF0 == 0xf0 {
			switch status {
			case 0xf0, 0xf7:
				event.message = SYSEX_MESSAGE

				length, err := binary.ReadUvarint(r)

				if err != nil {
					panic(err)
				}

				event.data = make([]byte, length)
				_, err = r.Read(event.data)

				if err != nil {
					panic(err)
				}
			case 0xff:
				event.message = META_MESSAGE
				metaStatus, err := r.ReadByte()
				if err != nil {
					panic(err)
				}
				event.status = metaStatus

				length, err := binary.ReadUvarint(r)

				if err != nil {
					panic(err)
				}

				event.data = make([]byte, length)
				_, err = r.Read(event.data)

				if err != nil && err != io.EOF {
					panic(err)
				}
			default:
				fmt.Println("Unknown event %#x\n", status)
			}
		} else {
			event.message = MIDI_MESSAGE

			switch status & 0xF0 {
				case 0x80, 0x90, 0xA0, 0xB0:
				event.data = make([]byte, 2)
				_, err := r.Read(event.data)

				if err != nil {
					panic(err)
				}
			case 0xC0, 0xD0:
				//midi short
				event.data = make([]byte, 1)
				_, err := r.Read(event.data)

				if err != nil {
					panic(err)
				}
			default:
				fmt.Println("Unknown event %#x\n", status)
			}
		}
	}
	return nil
}




func (event *IntermediateEvent) ProcessRawEvent () Event {
	switch event.message {
	case MIDI_MESSAGE:
		return ParseMidiEvent(event)
	case META_MESSAGE:
		return ParseMetaEvent(event)
	case SYSEX_MESSAGE:
		return ParseSysExEvent(event)
	default:
		fmt.Println("Unknown event %+v\n", event)
		return nil
	}
}
