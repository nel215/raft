package raft

import (
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func testRequestVote(t *testing.T) {
	cli, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	type testCase struct {
		Term     int64
		Expected bool
	}

	for _, c := range []testCase{
		testCase{Term: 0, Expected: true},
		testCase{Term: -1, Expected: false},
	} {
		args := &RequestVoteArgs{Term: c.Term}
		var reply RequestVoteResponse
		err = cli.Call("Service.RequestVote", args, &reply)

		if err != nil {
			t.Errorf("%s", err.Error())
		}

		if reply.VoteGranted != c.Expected {
			t.Errorf("VoteGranted is exptected %v, but got %v", c.Expected, reply.VoteGranted)
		}
	}
}

func setup(t *testing.T) {
	config := &Config{DBPath: "./test_db"}
	s, err := New(config)
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
}

func TestRPC(t *testing.T) {
	setup(t)
	testRequestVote(t)
}
