# OpenHue Go
![OpenHue Go Logo Medium Size](./docs/logo-md.png)

[![Build](https://github.com/openhue/openhue-go/actions/workflows/build.yml/badge.svg)](https://github.com/openhue/openhue-go/actions/workflows/build.yml)
[![Maintainability](https://api.codeclimate.com/v1/badges/ad99c96b6cbb59d2b81b/maintainability)](https://codeclimate.com/github/openhue/openhue-go/maintainability)
[![Go Reference](https://pkg.go.dev/badge/github.com/openhue/openhue-go.svg)](https://pkg.go.dev/github.com/openhue/openhue-go)

## Overview
OpenHue Go is a library written in Goland for interacting with the Philips Hue smart lighting systems.
This project is based on the [OpenHue API](https://github.com/openhue/openhue-api) specification. 
Therefore, most of its code is automatically generated thanks to the [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) project.

## Usage
Use the following command to import the library: 
```shell
go get -u github.com/openhue/openhue-go
```
And check the following example that toggles all the rooms of your house:
```go
package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"log"
)

func main() {

	home, _ := openhue.NewHome(openhue.LoadConfNoError())
	rooms, _ := home.GetRooms()

	for id, room := range rooms {

		fmt.Printf("> Toggling room %s (%s)\n", *room.Metadata.Name, id)

		for serviceId, serviceType := range room.GetServices() {

			if serviceType == openhue.ResourceIdentifierRtypeGroupedLight {
				groupedLight, _ := home.GetGroupedLightById(serviceId)

				home.UpdateGroupedLight(*groupedLight.Id, openhue.GroupedLightPut{
					On: groupedLight.Toggle(),
				})
			}
		}
	}
}
```
> [!NOTE]  
> The `openhue.LoadConf()` function allows loading the configuration from the well-known configuration file.
> Please refer to [this guide](https://www.openhue.io/cli/setup#manual-configuration) for more information.

### Bridge Discovery
Bridge Discovery on the local network has been made easy through the `BridgeDiscovery` helper: 
```go
package main

import (
	"fmt"
	"github.com/openhue/openhue-go"
	"log"
	"time"
)

func main() {

	bridge, err := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
	openhue.CheckErr(err)

	fmt.Println(bridge) // Output: Bridge{instance: "Hue Bridge - 1A3E4F", host: "ecb5fa1a3e4f.local.", ip: "192.168.1.xx"}
}
```
The `BridgeDiscovery.Discover()` function will first try to discover your local bridge via mDNS, 
and if that fails then it tries using [discovery.meethue.com](https://discovery.meethue.com) URL.

Options:
- `openhue.WithTimeout` allows setting the mDNS discovery timeout. Default value is `5` seconds.
- `openhue.WithDisabledUrlDiscovery` allows disabling the URL discovery.

### Authentication
Bridge authentication has been make simple via the `Authenticator` interface:
```go
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
```
In this example, we wait until the link button is pressed on the bridge. 
The `Authenticator.Authenticate()` function returns three values:
- `apiKey string` that is not empty when `retry = false` and `err == nil`
- `retry bool` which indicates that the link button has not been pressed
- `err error` which contains the error details

**You can consider the authentication has failed whenever the `retry` value is `false` and the `err` is not `nil`.**

## License
[![GitHub License](https://img.shields.io/github/license/openhue/openhue-cli)](https://github.com/openhue/openhue-cli/blob/main/LICENSE)

OpenHue is distributed under the [Apache License 2.0](http://www.apache.org/licenses/),
making it open and free for anyone to use and contribute to.
See the [license](./LICENSE) file for detailed terms.
