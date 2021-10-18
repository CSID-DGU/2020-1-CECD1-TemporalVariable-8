package mig

import (
	"encoding/json"
	"lightServer/ent"
)

type Session struct {
	Client *ent.User `json:"client"`
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
func (s *Session) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
