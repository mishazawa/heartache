package parser

import (
	"bytes"
	"encoding/binary"
)

type MidiFile struct {
	Format      uint16
	TracksCount uint16
	TimeDiv     uint16
	Tracks      []*Track
}

func ParseFile (data []byte) (*MidiFile, error) {
	r := bytes.NewReader(data)
	header, err := nextChunk(r)
	if err != nil {
		panic(err)
	}

	mfile := &MidiFile {
		Format:      binary.BigEndian.Uint16(header.Data[0:2]),
		TracksCount: binary.BigEndian.Uint16(header.Data[2:4]),
		TimeDiv:     binary.BigEndian.Uint16(header.Data[4:6]),
	}

	mfile.Tracks = make([]*Track, mfile.TracksCount)

	for i := range mfile.Tracks {
		mfile.Tracks[i] = NewTrack()
		trackData, err := nextChunk(r)
		if err != nil {
			panic(err)
		}
		mfile.Tracks[i].ParseEvents(trackData.Data)
	}
	return mfile, nil
}

type Chunk struct {
	Type string
	Data []byte
}

func nextChunk (r *bytes.Reader) (*Chunk, error) {
	buffer := make([]byte, 4)
	_, err := r.Read(buffer)
	if err != nil {
		return nil, err
	}

	chunkType := string(buffer)

	_, err = r.Read(buffer)
	if err != nil {
		return nil, err
	}

	rest := binary.BigEndian.Uint32(buffer)

	buffer = make([]byte, int(rest))

	_, err = r.Read(buffer)
	if err != nil {
		return nil, err
	}

	return &Chunk {chunkType, buffer}, nil
}
