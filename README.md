# OpenHue Go
[![Build](https://github.com/openhue/openhue-go/actions/workflows/build.yml/badge.svg)](https://github.com/openhue/openhue-go/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/openhue/openhue-go.svg)](https://pkg.go.dev/github.com/openhue/openhue-go)

## Overview

OpenHue Go is a library written in Goland for interacting with the Philips Hue smart lighting systems.

## Usage
Use the following command to import the library: 
```shell
go get github.com/openhue/openhue-go
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

	home, _ := openhue.NewHome(openhue.LoadConf())
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

## License
[![GitHub License](https://img.shields.io/github/license/openhue/openhue-cli)](https://github.com/openhue/openhue-cli/blob/main/LICENSE)

Open-Hue is distributed under the [Apache License 2.0](http://www.apache.org/licenses/),
making it open and free for anyone to use and contribute to.
See the [license](./LICENSE) file for detailed terms.