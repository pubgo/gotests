package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pubgo/xerror"
	"os"
	"path/filepath"
)

func NewWatcherDir(dir string) *WatcherDir {
	info, err := os.Lstat(dir)
	xerror.Panic(err)
	return &WatcherDir{name: dir, info: info}
}

type WatcherDir struct {
	name  string
	info  os.FileInfo
	files []os.FileInfo
	dirs  map[string]*WatcherDir
}

func (t *WatcherDir) Check() {

}

func main() {
	w := xerror.PanicErr(fsnotify.NewWatcher()).(*fsnotify.Watcher)
	go func() {
		for {
			select {
			case e, ok := <-w.Events:
				fmt.Println(ok, e.String())
			case err := <-w.Errors:
				fmt.Println(err)
				xerror.Panic(w.Close())
				return
			}
		}
	}()

	//d := xerror.PanicStr(filepath.Abs("."))
	d := xerror.PanicStr(filepath.Abs("/Users/barry/gopath/src/github.com/pubgo/gotests/watcher/watcher_dir/watcher.go"))
	fmt.Println(d)
	xerror.Panic(w.Add(d))
	select {}
}
