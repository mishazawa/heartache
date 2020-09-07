package parser

import (
	"fmt"
	"bytes"
	"encoding/binary"

	"github.com/mishazawa/heartache/parser/events"
)

type Track struct {
	Events []events.Event
}

func NewTrack () *Track {
	return &Track {
		Events: make([]events.Event, 0),
	}
}

func (t *Track) Parse (data []byte) {
	r := bytes.NewReader(data)
	for r.Len() != 0 {
		var event events.Event


		delta, err := binary.ReadUvarint(r)

		if err != nil {
			panic(err)
		}


		evt, err := r.ReadByte()
		if err != nil {
			panic(err)
		}

		// filter out all non-Midi events
		if evt & 0xf0 == 0xf0 {
			if evt == 0xf0 || evt == 0xf7 {
				// catch sysex events
				event, err = events.ParseSysExEvent(evt, r)
			} else if evt == 0xff {
				// catch meta events
				event, err = events.ParseMetaEvent(evt, r);
			} else {
				fmt.Printf("unknown META evt %x:%d\n", evt, delta)
			}
		} else {
			event, err = events.ParseMidiEvent(evt, r)
		}

		if event != nil {
			event.SetDelta(delta)
		}

		if err != nil {
			panic(err)
		}

		t.Events = append(t.Events, event)

	}
}
