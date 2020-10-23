package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/gen2brain/dlgs"
	"time"
)

func main() {
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)

	passwd, _, err := dlgs.Password("Password", "Enter your API key:")
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)

	yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		panic(err)
	}
	fmt.Println(yes)

	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}

	time.Sleep(time.Second)

	for {
		err = beeep.Notify("Title11", "Message \nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\nbodyMessage bodyMessage bodyMessage body\n", "assets/information.png")
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)

		err = beeep.Alert("Title22", "Message body", "assets/warning.png")
		if err != nil {
			panic(err)
		}
	}

}
