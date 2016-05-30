package raft

import (
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
	err = db.Put([]byte(CURRENT_TERM), []byte("0"), nil)
	if err != nil {
		return nil, err
	}
	return &Service{}, nil
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
	currentTerm, err := db.Get([]byte(CURRENT_TERM), nil)
	if err != nil {
		return err
	}
	fmt.Println(string(currentTerm))
	reply = &RequestVoteResponse{}

	return nil
}

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *string) error {
	return nil
}
