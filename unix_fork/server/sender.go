package main

import (
	"fmt"
	"github.com/pubgo/gotests/unix_fork"
	"golang.org/x/sys/unix"
	"net"
	"os"
	"syscall"
)

const (
	// 与receiver监听的unixsocks文件地址一致
	socksPath = "../unix_sock"
)

var oobSpace = unix.CmsgSpace(4)

func main() {
	file, err := os.OpenFile("./temp", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		panic(err)
	}
	fmt.Println(file.WriteString("hello"))
	//defer file.Close()

	fdnum := file.Fd()
	fmt.Println("fd", fdnum)
	fmt.Printf("%b %b %b %b\n", byte(fdnum), byte(fdnum>>8), byte(fdnum>>16), byte(fdnum>>24))
	fmt.Printf("ready to send fd: %d\n", fdnum)

	raddr, err := net.ResolveUnixAddr("unix", socksPath)
	if err != nil {
		panic(err)
	}
	// 连接UnixSock
	conn, err := net.DialUnix("unix", nil, raddr)
	if err != nil {
		panic(err)
	}

	//ln, _ := net.Listen("tcp", "localhost:9000")
	//rawConn, _ := ln.(net.Listener)

	//conn1, _ := ln.Accept()
	//conn1.()

	f, _ := conn.File()
	fdnum, _ = dupFd(fdnum)
	if err := unix_fork.SendFd(f, "temp", fdnum); err != nil {
		panic(err)
	}

	select {}
}

func dupFd(fd uintptr) (uintptr, error) {
	dupfd, _, errno := syscall.Syscall(syscall.SYS_FCNTL, fd, syscall.F_DUPFD_CLOEXEC, 0)
	if errno != 0 {
		return 0, fmt.Errorf("can't dup fd using fcntl: %s", errno)
	}
	return dupfd, nil
}
