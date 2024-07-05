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

```go
package main

import "github.com/openhue/openhue-go"
import "os"
import "log"

func main() {

	home, _ := openhue.NewHome(os.Getenv("BRIDGE"), os.Getenv("KEY"))
	lights, _ := home.GetLights()

	for id, light := range lights {
		fmt.Printf("Toggling light %s (%s)\n", *light.Metadata.Name, id)
		home.UpdateLight(*light.Id, openhue.LightPut{
			On: light.Toggle(),
		})
	}
}
```
This example demonstrates how to toggle all the lights of your house, in a very few lines of code.

## License
[![GitHub License](https://img.shields.io/github/license/openhue/openhue-cli)](https://github.com/openhue/openhue-cli/blob/main/LICENSE)

Open-Hue is distributed under the [Apache License 2.0](http://www.apache.org/licenses/),
making it open and free for anyone to use and contribute to.
See the [license](./LICENSE) file for detailed terms.