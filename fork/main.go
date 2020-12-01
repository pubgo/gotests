package main

import (
	"fmt"
	"github.com/pubgo/xerror"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var forkExec = func(argv0 string, argv []string) (pid int, err error) {
	return syscall.ForkExec(argv0, argv, &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
		Sys: &syscall.SysProcAttr{
			Setsid: true,
		},
	})
}

func main() {
	fmt.Println(os.Environ())
	fmt.Println()
	os.Setenv("test", "1111111")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-ch

	pid, err := forkExec(os.Args[0], os.Args)
	xerror.Panic(err)

	go func() {
		// 防止子进程变成僵尸进程
		for {
			_, _ = syscall.Wait4(pid, nil, syscall.WNOWAIT, nil)
			time.Sleep(time.Second * 5)
			return
		}
	}()
}
