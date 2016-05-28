package raft

import (
	"net/rpc"
	"testing"
)

func TestMain(t *testing.T) {
	s := new(Service)
	rpc.Register(s)
}
