package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

func main() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner
	for {
		fmt.Println("ssss")
		time.Sleep(time.Second)
	}
	//time.Sleep(4 * time.Second)                                 // Run for some time to simulate work
	//s.Stop()
}
