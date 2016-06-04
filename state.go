package raft

import (
	"bytes"
	"encoding/gob"
)

type State struct {
	CurrentTerm int64
	VotedFor    int64
}

func (s *State) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(s.CurrentTerm)
	if err != nil {
		return nil, err
	}
	err = encoder.Encode(s.VotedFor)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (s *State) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	err := decoder.Decode(&s.CurrentTerm)
	if err != nil {
		return err
	}
	return decoder.Decode(&s.VotedFor)
}
