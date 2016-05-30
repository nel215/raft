package raft

import (
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func TestMain(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	rpc.Register(s)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	go http.Serve(l, nil)

	cli, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		t.Errorf("%s", err.Error())
	}

	args := &RequestVoteArgs{}
	var reply RequestVoteResponse
	err = cli.Call("Service.RequestVote", args, &reply)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

}
