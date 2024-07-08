package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"log"
	"time"
)

func main() {

	bridge, err := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
	if err != nil {
		log.Fatal("Bridge Discovery Error: ", err)
	}

	fmt.Println(bridge)
}
