package main_test

import (
	"github.com/gen2brain/beeep"
	"testing"
)

func TestName(t *testing.T) {
	err := beeep.Alert("Title22", "Message body", "assets/warning.png")
	if err != nil {
		panic(err)
	}
}
