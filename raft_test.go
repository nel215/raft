package raft

import (
	"net/rpc"
	"testing"
)

func TestMain(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Errorf("%s", err)
	}
	rpc.Register(s)
}
