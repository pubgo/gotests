package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	f, err := os.OpenFile("filename.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			fmt.Println(f.WriteString("test\n"))
			time.Sleep(time.Second)
		}
	}()

	go func() {
		nfd, err := syscall.Dup(int(f.Fd()))
		if err != nil {
			panic(err)
		}
		f1 := os.NewFile(uintptr(nfd), "")
		for {
			var data = make([]byte, 1024)
			fmt.Println(f1.Read(data))
			fmt.Println(string(data))
			time.Sleep(time.Second)
			fmt.Println(f1.Seek(10,0))
		}
	}()

	select {}
}
