package redisutil

import (
	"bytes"
	"encoding/gob"
)

type GobBinary struct {
	Value interface{}
}

func (s *GobBinary) UnmarshalBinary(data []byte) error {
	return gob.NewDecoder(bytes.NewBuffer(data)).Decode(s.Value)
}

func (s *GobBinary) MarshalBinary() (data []byte, err error) {
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(s.Value); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

