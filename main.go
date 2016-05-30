package raft

import (
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
	db.Put([]byte(CURRENT_TERM), []byte("0"), nil)
	return &Service{}, nil
}

type RequestVoteArgs struct {
}
type AppendEntriesArgs struct{}

func (s *Service) RequestVote(args *RequestVoteArgs, reply *string) error {
	return nil
}

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *string) error {
	return nil
}
