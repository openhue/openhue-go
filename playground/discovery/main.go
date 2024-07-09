package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"time"
)

func main() {

	bridge, err := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
	openhue.CheckErr(err)

	fmt.Println(bridge)
}
