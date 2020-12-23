package main

import (
	"fmt"
	"github.com/pubgo/gotests/unix_fork"
	"net"
	"syscall"
	"time"
)

const (
	// socksPath unixsock文件所在地址
	socksPath = "../unix_sock"
)

func main() {
	// unlink删除已存在的unixSock文件
	syscall.Unlink(socksPath)
	laddr, err := net.ResolveUnixAddr("unix", socksPath)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenUnix("unix", laddr)
	if err != nil {
		panic(err)
	}

	//l.SetDeadline(time.Now().Add(time.Second * 10))
	fmt.Printf("waiting for conn from unix socks\n")
	conn, err := l.AcceptUnix()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	f, _ := conn.File()
	f1, err := unix_fork.RecvFd(f)
	if err != nil {
		panic(err)
	}

	for {
		// 从文件中读取文本内容
		buf := make([]byte, 1024)
		n, err := f1.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			//panic(err)
		}
		fmt.Printf("read %d data %s from file success\n", n, string(buf[:n]))
		time.Sleep(time.Second)
	}
}
