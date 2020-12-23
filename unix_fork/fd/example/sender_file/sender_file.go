package main

import (
	"flag"
	"fmt"
	"github.com/pubgo/gotests/unix_fork/fd"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"os"
	"time"
)

var (
	filename string
	socket   string
)

func init() {
	flag.StringVar(&filename, "f", "temp.txt", "filename")
	flag.StringVar(&socket, "s", "/tmp/sendfd.sock", "socket")
}

func main() {
	flag.Parse()

	if !flag.Parsed() || filename == "" || socket == "" {
		flag.Usage()
		os.Exit(1)
	}

	unix.Unlink(socket)

	fmt.Println(filename)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	//dt, err := ioutil.ReadAll(f)
	//fmt.Println(string(dt), err)

	go func() {
		for {
			fmt.Println(f.Write([]byte("sender\n")))
			fmt.Println(f.Sync())
			time.Sleep(time.Second)
			//dt, err := ioutil.ReadFile(filename)
			//fmt.Println(string(dt), err)
		}
	}()

	l, err := net.Listen("unix", socket)
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
}
