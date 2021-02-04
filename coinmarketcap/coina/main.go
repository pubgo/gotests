package main

import (
	"fmt"
	"net/http"
	"time"

	coingecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	cg := coingecko.NewClient(httpClient)
	fmt.Println(cg.Ping())

	// fmt.Println(cg.CoinsList())
	fmt.Println(cg.)
}
