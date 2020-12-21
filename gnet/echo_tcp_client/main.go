package main

import (
	"fmt"
	"net"
	"time"
)

// Example command: go run client.go
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		conn.Write([]byte("hello"))
		time.Sleep(time.Millisecond * 100)
		var data = make([]byte, 1024)
		conn.Read(data)
		fmt.Println(string(data))
	}
}
