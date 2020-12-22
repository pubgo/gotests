package main

import (
	"fmt"
	"github.com/cloudflare/tableflip"
	"github.com/pubgo/xerror"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Example command: go run client.go
func main() {
	upg, _ := tableflip.New(tableflip.Options{})
	defer upg.Stop()

	ln, _ := upg.Listen("tcp", "localhost:9000")
	defer ln.Close()

	go func() {
		for i := 0; ; i++ {
			conn, err := upg.Conn("tcp", fmt.Sprintf("%d", i))
			xerror.Panic(err)
			if conn == nil {
				break
			}
			ioutil.WriteFile("test111.txt", []byte(fmt.Sprintf("%d", i)), 0755)
			fmt.Println(i)

			go func() {
				for {
					conn.Write([]byte("hello"))
					time.Sleep(time.Millisecond * 500)
					var data = make([]byte, 1024)
					conn.Read(data)
					fmt.Println(string(data))
				}
			}()
		}
	}()

	go func() {
		for i := 0; ; i++ {
			conn, err := ln.Accept()
			if err != nil {
				break
			}

			xerror.Panic(upg.AddConn("tcp", fmt.Sprintf("%d", i), conn.(tableflip.Conn)))

			go func() {
				for {
					conn.Write([]byte("hello"))
					time.Sleep(time.Millisecond * 500)
					var data = make([]byte, 1024)
					conn.Read(data)
					fmt.Println(string(data))
				}
			}()
		}
	}()

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
		<-ch
		xerror.Panic(upg.Upgrade())
	}()

	if err := upg.Ready(); err != nil {
		panic(err)
	}

	<-upg.Exit()
	time.Sleep(time.Second * 5)
}
