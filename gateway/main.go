package main

import (
	"fmt"
	_ "github.com/jackpal/gateway"
)

func main() {
	fmt.Println(gateway.DiscoverGateway())
}
