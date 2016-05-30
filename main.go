package raft

import (
	"bytes"
	"encoding/gob"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	CURRENT_TERM = "CURRENT_TERM"
)

type Service struct {
	config *Config
}

type Config struct {
	DBPath string
}

func (s *Service) InitDB() error {
	db, err := leveldb.OpenFile(s.config.DBPath, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	enc.Encode(&State{CurrentTerm: 0})

	return db.Put([]byte(CURRENT_TERM), buf.Bytes(), nil)
}

func New(config *Config) (*Service, error) {
	s := &Service{config: config}
	err := s.InitDB()
	if err != nil {
		return nil, err
	}
	return s, nil
}

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

type RequestVoteArgs struct {
	Term int64
}
type RequestVoteResponse struct {
	VoteGranted bool
}
type AppendEntriesArgs struct{}

func (s *Service) RequestVote(args *RequestVoteArgs, reply *RequestVoteResponse) error {
	db, err := leveldb.OpenFile(s.config.DBPath, nil)
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
	if err != nil {
		return err
	}

	*reply = RequestVoteResponse{VoteGranted: false}
	if args.Term < state.CurrentTerm {
		return nil
	}

	// TODO: and candicate's log is at least as up-to-date as receiver's log
	if state.VotedFor == 0 {
		reply.VoteGranted = true
	}

	return nil
}

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *string) error {
	return nil
}
