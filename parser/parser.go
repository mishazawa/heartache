package parser

import (
	"bytes"
	"io/ioutil"
	"encoding/binary"
)

type MidiFile struct {
	Format      uint16
	TracksCount uint16
	TimeDiv     uint16
	Tracks      []*Track
}

func ParseFile (name string) (*MidiFile, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return parseData(file)
}

func parseData (data []byte) (*MidiFile, error) {
	r := bytes.NewReader(data)
	header, err := nextChunk(r)

	if err != nil {
		return nil, err
	}

	mfile := &MidiFile {
		Format:      binary.BigEndian.Uint16(header[0:2]),
		TracksCount: binary.BigEndian.Uint16(header[2:4]),
		TimeDiv:     binary.BigEndian.Uint16(header[4:6]),
	}

	mfile.Tracks = make([]*Track, mfile.TracksCount)


	for i := range mfile.Tracks {
		track := newTrack()

		data, err := nextChunk(r)
		if err != nil {
			return nil, err
		}

		err = track.parseEvents(data)
		if err != nil {
			return nil, err
		}

		mfile.Tracks[i] = track
	}

	return mfile, nil
}



