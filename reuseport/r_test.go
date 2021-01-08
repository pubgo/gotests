package reuseport

import (
	"testing"
)

func TestName(t *testing.T) {
	id, addr, err := TCPSocket("tcp", ":8082", true)
	if err != nil {
		panic(err)
	}


}
