<div align="center">

# OpenHue Go

![OpenHue Go Logo](./docs/logo-md.png)

**A modern Go library for Philips Hue smart lighting systems**

[![Build](https://github.com/openhue/openhue-go/actions/workflows/build.yml/badge.svg)](https://github.com/openhue/openhue-go/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/openhue/openhue-go)](https://goreportcard.com/report/github.com/openhue/openhue-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/openhue/openhue-go.svg)](https://pkg.go.dev/github.com/openhue/openhue-go)
[![GitHub License](https://img.shields.io/github/license/openhue/openhue-go)](https://github.com/openhue/openhue-go/blob/main/LICENSE)

[Features](#features) • [Installation](#installation) • [Quick Start](#quick-start) • [Documentation](#documentation) • [Contributing](#contributing)

</div>

---

## Features

🏠 **Home Abstraction** — Intuitive entry point to interact with your entire Hue ecosystem  
🔍 **Bridge Discovery** — Automatic bridge detection via mDNS with URL fallback  
🔐 **Easy Authentication** — Simplified link button authentication flow  
💡 **Full API Coverage** — Auto-generated from the [OpenHue API](https://github.com/openhue/openhue-api) specification  
⚡ **Type-Safe** — Strongly typed Go structs for all Hue resources  

## Installation

**Requirements:** Go 1.23+

```shell
go get -u github.com/openhue/openhue-go
```

## Quick Start

Toggle all rooms in your home with just a few lines of code:

```go
package main

import (
    "fmt"
    "github.com/openhue/openhue-go"
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
> The `openhue.LoadConf()` function loads configuration from the well-known configuration file.  
> See the [configuration guide](https://www.openhue.io/cli/setup#manual-configuration) for details.

## Documentation

### Bridge Discovery

Automatically discover Hue bridges on your local network:

```go
bridge, err := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
openhue.CheckErr(err)

fmt.Println(bridge)
// Output: Bridge{instance: "Hue Bridge - 1A3E4F", host: "ecb5fa1a3e4f.local.", ip: "192.168.1.xx"}
```

The discovery process tries mDNS first, then falls back to [discovery.meethue.com](https://discovery.meethue.com).

**Options:**
- `openhue.WithTimeout(duration)` — Set mDNS discovery timeout (default: 5 seconds)
- `openhue.WithDisabledUrlDiscovery` — Disable URL fallback discovery

### Authentication

Authenticate with your bridge using the link button:

```go
bridge, _ := openhue.NewBridgeDiscovery(openhue.WithTimeout(1 * time.Second)).Discover()
authenticator, _ := openhue.NewAuthenticator(bridge.IpAddress)

fmt.Println("Press the link button")

var key string
for len(key) == 0 {
    apiKey, retry, err := authenticator.Authenticate()

    if err != nil && retry {
        fmt.Printf(".")
        time.Sleep(500 * time.Millisecond)
    } else if err != nil && !retry {
        openhue.CheckErr(err)
    } else {
        key = apiKey
    }
}

fmt.Println("\n", key)
```

The `Authenticate()` function returns:
- `apiKey` — The API key (non-empty on success)
- `retry` — `true` if the link button hasn't been pressed yet
- `err` — Error details if authentication failed

> [!TIP]
> Authentication has failed when `retry == false` and `err != nil`.

## Related Projects

- [OpenHue API](https://github.com/openhue/openhue-api) — OpenAPI specification for Philips Hue
- [OpenHue CLI](https://github.com/openhue/openhue-cli) — Command-line interface built with this library
- [openhue.io](https://www.openhue.io/) — Official documentation and guides

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.

This project uses [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) for code generation. Most code in `openhue.gen.go` is auto-generated — see the project structure before making changes.

## License

OpenHue Go is distributed under the [Apache License 2.0](http://www.apache.org/licenses/), making it open and free for anyone to use and contribute to. See the [LICENSE](./LICENSE) file for details.
