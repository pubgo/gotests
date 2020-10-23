package main

import (
	"fmt"
	"github.com/pubgo/xerror"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

// clipboard represents the system clipboard
type clipboard struct {
	window *glfw.Window
}

// Content returns the clipboard content
func (c *clipboard) Content() string {
	return glfw.GetClipboardString()
}

// SetContent sets the clipboard content
func (c *clipboard) SetContent(content string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r.(error).Error() + "ssss")
		}
	}()

	glfw.SetClipboardString(content)
}

func main() {
	xerror.Exit(glfw.Init())

	c := &clipboard{}
	go func() {
		for range time.Tick(time.Second) {
			c.SetContent(time.Now().String())
		}
	}()

	for range time.Tick(time.Second) {
		text := c.Content()
		fmt.Println(text)
	}
}
