package main_test

import (
	"github.com/gen2brain/beeep"
	"testing"
)

func TestName(t *testing.T) {
	err := beeep.Alert("Title22", "https://dev.to/t/go \n Message body", "./assets/warning.png")
	if err != nil {
		panic(err)
	}
}

func TestNotify(t *testing.T) {
	err := beeep.Notify("Title11", "https://dev.to/t/go \n Message \nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\n", "assets/information.png")
	if err != nil {
		panic(err)
	}
}

func TestBeep(t *testing.T) {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}
