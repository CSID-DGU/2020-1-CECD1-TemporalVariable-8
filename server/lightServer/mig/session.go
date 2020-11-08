package mig

import (
	"encoding/json"
	"golang.org/x/oauth2"
)

type Session struct {
	Token *oauth2.Token `json:"token"`
}

func (s *Session) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}
func (s *Session) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
