package raft

type Service struct {
}

type RequestVoteArgs struct{}
type AppendEntriesArgs struct{}

func (s *Service) RequestVote(args *RequestVoteArgs, reply *string) error {
	return nil
}

func (s *Service) AppendEntries(args *AppendEntriesArgs, reply *string) error {
	return nil
}
