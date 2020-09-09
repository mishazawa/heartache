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

func (t *Track) ParseEvents (data []byte) {
	rawEvents := make([]*events.IntermediateEvent, 0)

	r := bytes.NewReader(data)
	var runningStatus byte

	for r.Len() > 0 {
		var err           error
		var delta         uint64
		var nextByte      byte

		intermediateEvent := events.NewIntermediateEvent();

		delta, err = binary.ReadUvarint(r)

		if err != nil {
			panic(err)
		}

		nextByte, err = r.ReadByte()

		if err != nil {
			panic(err)
		}

		kind := nextByte & 0xf0
		buffer := make([]byte, 0)

		if kind >= 0x80 {
			runningStatus = nextByte
		} else {
			buffer = append(buffer, nextByte)
		}

		err = intermediateEvent.ParseEvent(runningStatus, buffer, delta, r)

		if err != nil {
			panic(err)
		}

		rawEvents = append(rawEvents, intermediateEvent)
	}

	t.Events = make([]events.Event, len(rawEvents))

	for i, ev := range rawEvents {
		t.Events[i] = ev.ProcessRawEvent()
		fmt.Printf("%T\n", t.Events[i])
	}

}

