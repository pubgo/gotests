package main

import (
	"encoding/json"
	"fmt"
	"github.com/pubgo/xerror"
	"log"

	gecko "github.com/superoo7/go-gecko/v3"
)

func main() {
	cg := gecko.NewClient(nil)
	tickers, err := cg.CoinsIDTickers("bitcoin", 0)
	if err != nil {
		log.Fatal(err)
	}

	dt, err := json.MarshalIndent(cg, "", "\t")
	xerror.Panic(err)
	fmt.Println(string(dt))

	fmt.Println(len(tickers.Tickers))
	tickers, err = cg.CoinsIDTickers("bitcoin", 1)
	fmt.Println(len(tickers.Tickers))

	dt, err = json.MarshalIndent(cg, "", "\t")
	xerror.Panic(err)
	fmt.Println(string(dt))
}
