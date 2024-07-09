package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"time"
)

func main() {

	bridge, err := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
	openhue.CheckErr(err)

	authenticator, err := openhue.NewAuthenticator(bridge.IpAddress)
	openhue.CheckErr(err)

	fmt.Println("Press the link button")

	var key string
	for len(key) == 0 {

		// try to authenticate
		apiKey, retry, err := authenticator.Authenticate()

		if err != nil && retry {
			// link button not pressed
			fmt.Printf(".")
			time.Sleep(500 * time.Millisecond)
		} else if err != nil && !retry {
			// there is a real error
			openhue.CheckErr(err)
		} else {
			key = apiKey
		}
	}

	fmt.Println("\n", key)
}
