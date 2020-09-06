package parser

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type Track struct {
	Events []Event
}

func NewTrack () *Track {
	return &Track {
		Events: make([]Event, 0),
	}
}

func (t *Track) Parse (data []byte) {
	r := bytes.NewReader(data)
	for r.Len() > 0 {
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
				event, _ := ParseSysExEvent(delta, evt, r) // now it panic
				t.Events = append(t.Events, event)

			} else if evt == 0xff {
				// catch meta events
				event, _ := ParseMetaEvent(delta, evt, r);
				t.Events = append(t.Events, event)
			} else {
				fmt.Printf("unknown META evt %x:%d\n", evt, delta)
				t.Events = append(t.Events, nil)
			}
		} else {
			// fmt.Printf("unknown MIDI %x:%d\n", evt, delta)
			t.Events = append(t.Events, nil)
		}
	}
}
