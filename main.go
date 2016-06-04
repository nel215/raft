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

type RequestVoteArgs struct {
	Term int64
}
type RequestVoteResponse struct {
	VoteGranted bool
}
type AppendEntriesArgs struct {
	Term int64
}
type AppendEntriesResponse struct {
	Term    int64
	Success bool
}

func (s *Service) getState() (*State, error) {
	db, err := leveldb.OpenFile(s.config.DBPath, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	data, err := db.Get([]byte(CURRENT_TERM), nil)
	if err != nil {
		return nil, err
	}
	state := new(State)
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(state)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func (s *Service) RequestVote(args *RequestVoteArgs, reply *RequestVoteResponse) error {
	state, err := s.getState()
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

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *AppendEntriesResponse) error {
	state, err := s.getState()
	if err != nil {
		return err
	}
	*reply = AppendEntriesResponse{Term: state.CurrentTerm, Success: false}
	// $5.1
	if args.Term < state.CurrentTerm {
		return nil
	}
	return nil
}
