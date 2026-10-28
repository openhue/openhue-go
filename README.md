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

🏠 **Home Abstraction** — One entry point for lights, rooms, zones, scenes, sensors, and more  
🔍 **Bridge Discovery** — Automatic bridge detection via mDNS with URL fallback  
🔐 **Easy Authentication** — Simplified link button authentication flow  
💡 **Full API Coverage** — Generated from the [OpenHue API](https://github.com/openhue/openhue-api) specification  
⚡ **Type-Safe** — Strongly typed Go structs for all Hue resources  
🔒 **Secure by Default** — Bridge TLS certificates verified against the embedded Hue root CAs  
🧪 **Testable** — Built-in mock via `NewTestHome()` for unit-testing your own code  

## Installation

**Requirements:** Go 1.24+ and a Philips Hue bridge reachable on your network

```shell
go get github.com/openhue/openhue-go
```

## Quick Start

Toggle all rooms in your home with just a few lines of code:

```go
package main

import (
    "context"
    "fmt"

    "github.com/openhue/openhue-go"
)

func main() {
    home, err := openhue.NewHome(openhue.LoadConfNoError())
    openhue.CheckErr(err)

    ctx := context.Background()
    rooms, err := home.GetRooms(ctx)
    openhue.CheckErr(err)

    for id, room := range rooms {
        fmt.Printf("> Toggling room %s (%s)\n", room.Metadata.Name, id)

        for serviceId, serviceType := range room.GetServices() {
            if serviceType == openhue.ResourceIdentifierRtypeGroupedLight {
                groupedLight, err := home.GetGroupedLightById(ctx, serviceId)
                openhue.CheckErr(err)

                err = home.UpdateGroupedLight(ctx, groupedLight.Id, openhue.GroupedLightPut{
                    On: groupedLight.Toggle(),
                })
                openhue.CheckErr(err)
            }
        }
    }
}
```

> [!NOTE]
> `openhue.LoadConfNoError()` returns the bridge IP and API key read from the well-known configuration file
> (use `openhue.LoadConf()` if you prefer to handle the error yourself).
> See the [configuration guide](https://www.openhue.io/cli/setup#manual-configuration) for details.
> No config yet? Follow [Bridge Discovery](#bridge-discovery) and [Authentication](#authentication) below, then pass the values to `openhue.NewHome(ip, apiKey)`.

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
- `openhue.WithDisabledUrlDiscovery()` — Disable URL fallback discovery

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

Options for `NewAuthenticator`: `openhue.WithDeviceType(name)` and `openhue.WithGenerateClientKey(bool)`.

The `Authenticate()` function returns:
- `apiKey` — The API key (non-empty on success)
- `retry` — `true` if the link button hasn't been pressed yet
- `err` — Error details if authentication failed

> [!TIP]
> Authentication has failed for good when `retry == false` and `err != nil`.

### Creating a Home

```go
home, err := openhue.NewHome("192.168.1.10", apiKey,
    openhue.WithRequestTimeout(10*time.Second),
)
```

**Options:**
- `openhue.WithRequestTimeout(duration)` — Set the timeout of requests sent to the bridge
- `openhue.WithCustomHTTPClient(client)` — Use your own `*http.Client` (for example, to customize TLS or transport)

> [!NOTE]
> By default, the bridge certificate is verified against the embedded Hue root CAs. `InsecureSkipVerify` is never used.

### Controlling Resources

Every method takes a `context.Context`. List methods return a `map[id]Resource`, and `Get...ById` methods return a pointer to a single resource.

| Resource | Methods |
|----------|---------|
| Lights | `GetLights`, `GetLightById`, `UpdateLight` |
| Grouped lights | `GetGroupedLights`, `GetGroupedLightById`, `UpdateGroupedLight` |
| Rooms | `GetRooms`, `GetRoomById`, `CreateRoom`, `UpdateRoom`, `DeleteRoom` |
| Zones | `GetZones`, `GetZoneById`, `CreateZone`, `UpdateZone`, `DeleteZone` |
| Scenes | `GetScenes`, `GetSceneById`, `CreateScene`, `UpdateScene`, `DeleteScene`, `ActivateScene` |
| Smart scenes | `GetSmartScenes`, `GetSmartSceneById`, `CreateSmartScene`, `UpdateSmartScene`, `DeleteSmartScene` |
| Devices | `GetDevices`, `GetDeviceById`, `GetDevicePowers`, `GetDevicePowerById` |
| Sensors & buttons | `GetMotionSensors`, `GetTemperatureSensors`, `GetButtons` (and `...ById`) |
| Entertainment | `GetEntertainmentConfigurations`, `StartEntertainment`, `StopEntertainment`, `UpdateEntertainmentConfiguration` |
| Bridge | `GetBridge`, `GetBridges`, `UpdateBridge`, `GetBridgeHome`, `GetResources` |

Example — list lights and toggle one:

```go
lights, err := home.GetLights(ctx)
openhue.CheckErr(err)

for id, light := range lights {
    fmt.Printf("%s: %s\n", id, light.Metadata.Name)

    err = home.UpdateLight(ctx, id, openhue.LightPut{On: light.Toggle()})
    openhue.CheckErr(err)
}
```

> [!IMPORTANT]
> Generated types follow the OpenAPI spec: required fields (such as `Id`, `Metadata.Name`, or `On.On`) are plain values, not pointers.
> If you are upgrading from an earlier version, drop the `*` dereferences (`room.Metadata.Name` instead of `*room.Metadata.Name`).

### Error Handling

Failed API calls return an `*openhue.ApiError` that works with `errors.Is`:

```go
_, err := home.GetLightById(ctx, "unknown-id")
if errors.Is(err, openhue.ErrNotFound) {
    fmt.Println("no such light")
}
```

Available sentinels: `ErrUnauthorized`, `ErrForbidden`, `ErrNotFound`, `ErrConflict`, `ErrTooManyRequests`, `ErrInternalServerError`, `ErrServiceUnavailable`, `ErrEmptyResponse`.

### Testing Your Code

`openhue.NewTestHome()` returns a `Home` backed by a [testify](https://github.com/stretchr/testify) mock of the generated client, so you can unit-test code that uses this library without a bridge:

```go
home, apiMock := openhue.NewTestHome()
// configure expectations on apiMock, then pass `home` to the code under test
```

### Examples

Runnable examples live in [`playground/`](./playground) (`authenticate`, `discovery`, `list_lights`, `toggle_lights`, `toggle_rooms`). They require a real bridge.

```shell
go run ./playground/list_lights
```

## Related Projects

- [OpenHue API](https://github.com/openhue/openhue-api) — OpenAPI specification for Philips Hue
- [OpenHue CLI](https://github.com/openhue/openhue-cli) — Command-line interface built with this library
- [openhue.io](https://www.openhue.io/) — Official documentation and guides

## Contributing

Contributions are welcome! Feel free to open issues and pull requests.

`openhue.gen.go` is generated by [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) and must **not** be edited by hand. Add features in the hand-written files (`openhue.go`, `helpers.go`, ...).

```shell
make generate   # regenerate openhue.gen.go from the latest spec (or: make generate spec=./openhue.yaml)
make test       # run the tests with coverage
make tidy       # tidy go.mod
```

Commits follow [Conventional Commits](https://www.conventionalcommits.org/) (`feat:`, `fix:`, `docs:`, ...). See [AGENTS.md](./AGENTS.md) for more details on the code layout.

## License

OpenHue Go is distributed under the [Apache License 2.0](http://www.apache.org/licenses/), making it open and free for anyone to use and contribute to. See the [LICENSE](./LICENSE) file for details.
