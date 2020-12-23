package main

import (
	"flag"
	"fmt"
	"github.com/pubgo/gotests/unix_fork/fd"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	port   int
	socket string
)

func init() {
	flag.IntVar(&port, "p", 9000, "listen port")
	flag.StringVar(&socket, "s", "/tmp/sendfd.sock", "socket")
}

func main() {
	flag.Parse()

	if !flag.Parsed() || socket == "" {
		flag.Usage()
		os.Exit(1)
	}

	unix.Unlink(socket)

	tcpl, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer tcpl.Close()

	var c net.Conn
	c, err = tcpl.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	go func() {
		for {
			var data = make([]byte, 1024)
			c.Read(data)
			fmt.Println(string(data))
			c.Write([]byte("sender"))
			time.Sleep(time.Second)
		}
	}()

	var f *os.File
	f, err = c.(*net.TCPConn).File()
	if err != nil {
		log.Fatal(err)
	}

	var l net.Listener
	l, err = net.Listen("unix", socket)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	var a net.Conn
	a, err = l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	listenConn := a.(*net.UnixConn)
	if err = fd.Put(listenConn, f); err != nil {
		log.Fatal(err)
	}
	select {}
}
