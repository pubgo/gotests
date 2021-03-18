package main

import (
	"fmt"
	reuse "github.com/libp2p/go-reuseport"
	"net/http"
	"time"
)

func main() {
	// listen on the same port. oh yeah.
	l1, err := reuse.Listen("tcp", "127.0.0.1:1234")
	fmt.Println(err)
	m1 := http.NewServeMux()
	m1.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("l1", *request)
		writer.Write([]byte("ok"))
	})
	go http.Serve(l1, m1)

	l2, err := reuse.Listen("tcp", "127.0.0.1:1234")
	fmt.Println(err)

	for i := 0; ; i++ {
		time.Sleep(time.Second * 5)

		m2 := http.NewServeMux()
		m2.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Println(fmt.Sprint("l",i), *request)
			writer.Write([]byte("ok"))
		})
		go http.Serve(l2, m2)
	}

	select {}

	// dial from the same port. oh yeah.
	//l1, _ := reuse.Listen("tcp", "127.0.0.1:1234")
	//l2, _ := reuse.Listen("tcp", "127.0.0.1:1235")
	//c, _ := reuse.Dial("tcp", "127.0.0.1:1234", "127.0.0.1:1235")
}
