package parser

import (
	"bytes"
	"encoding/binary"
)

type chunk []byte;

func nextChunk (r *bytes.Reader) (chunk, error) {
	buffer := make([]byte, 4)

	_, err := r.Read(buffer)
	if err != nil {
		return nil, err
	}

	_, err = r.Read(buffer)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint32(buffer)

	buffer = make([]byte, int(length))

	_, err = r.Read(buffer)
	if err != nil {
		return nil, err
	}

	return chunk(buffer), nil
}
