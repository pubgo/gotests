package main

import (
	"fmt"
	"os"

	"github.com/anaskhan96/soup"
)

func main() {
	resp, err := soup.Get("https://xkcd.com")
	if err != nil {
		os.Exit(1)
	}

	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}
