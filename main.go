package raft

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	CURRENT_TERM = "CURRENT_TERM"
)

type Service struct {
}

func New() (*Service, error) {
	db, err := leveldb.OpenFile("./test_db", nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(&State{CurrentTerm: 0})

	err = db.Put([]byte(CURRENT_TERM), buf.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return &Service{}, nil
}

type State struct {
	CurrentTerm int64
}

func (s *State) GobEncode() ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	err := encoder.Encode(s.CurrentTerm)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func (s *State) GobDecode(buf []byte) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)
	return decoder.Decode(&s.currentTerm)
}

type RequestVoteArgs struct {
}
type RequestVoteResponse struct {
}
type AppendEntriesArgs struct{}

func (s *Service) RequestVote(args *RequestVoteArgs, reply *RequestVoteResponse) error {
	db, err := leveldb.OpenFile("./test_db", nil)
	if err != nil {
		return err
	}
	defer db.Close()
	data, err := db.Get([]byte(CURRENT_TERM), nil)
	if err != nil {
		return err
	}
	state := new(State)
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(state)
	fmt.Println(state)
	reply = &RequestVoteResponse{}

	return nil
}

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *string) error {
	return nil
}
