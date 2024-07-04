# OpenHue Go
[![Build](https://github.com/openhue/openhue-go/actions/workflows/build.yml/badge.svg)](https://github.com/openhue/openhue-go/actions/workflows/build.yml)

## Overview

OpenHue Go is a library written in Goland for interacting with the Philips Hue smart lighting systems.

## Usage

```go
package main

import "os"
import "github.com/openhue/openhue-go"
import "log"

func main() {

	home, err := openhue.NewHome(os.Getenv("BRIDGE"), os.Getenv("KEY"))
	if err != nil {
		log.Fatal(err)
	}

	lights, err := home.GetLights()
	if err != nil {
		log.Fatal(err)
	}

	for _, light := range lights {

		body := openhue.UpdateLightJSONRequestBody{
			On: light.On.Switch(),
		}
		home.SetLight(*light.Id, body)
	}
}
```

## License
[![GitHub License](https://img.shields.io/github/license/openhue/openhue-cli)](https://github.com/openhue/openhue-cli/blob/main/LICENSE)

Open-Hue is distributed under the [Apache License 2.0](http://www.apache.org/licenses/),
making it open and free for anyone to use and contribute to.
See the [license](./LICENSE) file for detailed terms.