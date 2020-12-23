package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

var forkExec = func(argv0 string, argv []string, files ...uintptr) (pid int, err error) {
	lll := []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}
	lll = append(lll, files...)
	fmt.Println(lll)
	return syscall.ForkExec(argv0, argv, &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: lll,
		Sys: &syscall.SysProcAttr{
			Setsid: true,
		},
	})
}

// Example command: go run client.go
func main() {
	var connList []uintptr
	isUpgrade := os.Getenv("fork") != ""
	var err error
	var ln net.Listener
	if isUpgrade {
		ln, err = net.FileListener(os.NewFile(3, ""))
		if err != nil {
			panic(err)
		}
	} else {
		ln, _ = net.Listen("tcp", "localhost:9000")
	}

	if isUpgrade {
		conn, err := net.FileConn(os.NewFile(4, ""))
		Panic(err)

		go func() {
			for {
				conn.Write([]byte("hello" + os.Getenv("fork")))
				time.Sleep(time.Millisecond * 500)
				var data = make([]byte, 1024)
				conn.Read(data)
				fmt.Println(string(data))
			}
		}()
	}

	//defer ln.Close()

	rawConn, ok := ln.(syscall.Conn)
	if !ok {
		panic("not raw")
	}

	raw, err := rawConn.SyscallConn()
	Panic(err)

	var dupfd uintptr
	Panic(raw.Control(func(fd uintptr) {
		dupfd, err = dupFd(fd)
		if err != nil {
			panic(err)
		}
		Panic(err)
	}))

	connList = append(connList, dupfd)

	go func() {
		for i := 0; ; i++ {
			conn, err := ln.Accept()
			if err != nil {
				break
			}

			rawConn, ok := conn.(syscall.Conn)
			if !ok {
				panic("not raw")
			}

			raw, err := rawConn.SyscallConn()
			Panic(err)

			var dupfd uintptr
			Panic(raw.Control(func(fd uintptr) {
				dupfd, err = dupFd(fd)
				Panic(err)
			}))

			connList = append(connList, dupfd)

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

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-ch

	os.Setenv("fork", "true")

	pid, err := forkExec(os.Args[0], os.Args, connList...)
	Panic(err)

	go func() {
		// 防止子进程变成僵尸进程
		for {
			_, _ = syscall.Wait4(pid, nil, syscall.WNOWAIT, nil)
			time.Sleep(time.Second * 5)
			return
		}
	}()

	time.Sleep(time.Second * 5)
}

func dupFd(fd uintptr) (uintptr, error) {
	dupfd, _, errno := syscall.Syscall(syscall.SYS_FCNTL, fd, syscall.F_DUPFD_CLOEXEC, 0)
	if errno != 0 {
		return 0, fmt.Errorf("can't dup fd using fcntl: %s", errno)
	}
	return dupfd, nil
}
