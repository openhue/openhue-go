# OpenHue Go Library - Code Review & Improvement Suggestions

## Executive Summary
This is a well-structured library with a clean foundation. The auto-generation approach is solid, and the high-level wrapper pattern is appropriate. However, there are several critical issues and many opportunities to enhance the library for production use.

---

## 🔴 Critical Issues

### 1. **Global HTTP Transport Mutation (SECURITY RISK)** ✅ **FIXED**
**Status**: Fixed in branch `fix/secure-tls-verification` (commit d4669f8)

**Original Location**: `openhue.go:304`
```go
http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
```

**Problem**: This modifies the global `http.DefaultTransport`, affecting **all** HTTP clients in the entire application, not just this library.

**Impact**: 
- Any other HTTP client in the user's application will have TLS verification disabled
- This is a **major **Solution Implemented**:
- Removed global `http.DefaultTransport` mutation
- Created dedicated HTTP client with custom transport for each `Home` instance
- Added Philips Hue Bridge root CA certificates (`certs.go`)
- Implemented proper certificate chain verification via `VerifyPeerCertificate` callback
- Supports IP address-based connections (certificates don't have IP SANs)
- Certificates validated against trusted Hue root CAs
- Only hostname/IP validation is skipped (necessary for local IoT devices)

### 2. **Missing Context Propagation** ✅ **FIXED**
**Status**: Fixed in branch `feat/context-propagation` (commit e10a4c6)
**Locations**: Throughout `openhue.go`

**Problem**: All methods use `context.Background()` instead of accepting context from callers.

**Impact**:
- Cannot cancel operations
- Cannot set timeouts
- Cannot propagate request-scoped values
- Not compatible with modern Go HTTP patterns

**Solution Implemented**:
- All 13 public methods now accept `context.Context` as first parameter
- Users can now cancel operations, set timeouts, and propagate request-scoped values
- Fixed error handling in `UpdateLight` and `UpdateGroupedLight` (now check HTTP status codes)
- Updated all example code and tests to use context

**Breaking Change**: Users must update their code to pass context:
```go
// Before: home.GetLights()
// After:  home.GetLights(context.Background())
```

### 3. **Error Response Handling is Incomplete** ⚠️ **PARTIALLY FIXED**
**Status**: Fixed for `UpdateLight` and `UpdateGroupedLight` in branch `feat/context-propagation`

**Location**: `openhue.go` - methods like `UpdateLight`

**Problem**: Some update methods don't check HTTP status codes:
```go
func (h *Home) UpdateLight(lightId string, body LightPut) error {
    _, err := h.api.UpdateLightWithResponse(context.Background(), lightId, body)
    return err // No status code check!
}
```

**Solution**: Always check response status codes like in `UpdateScene`:
```go
func (h *Home) UpdateLight(lightId string, body LightPut) error {
    resp, err := h.api.UpdateLightWithResponse(context.Background(), lightId, body)
    if err != nil {
        return err
    }
    if resp.HTTPResponse.StatusCode != http.StatusOK {
        return newApiError(resp)
    }
    return nil
}
```

### 4. **Unsafe Array Indexing**
**Locations**: `openhue.go:50`, `openhue.go:112`, `openhue.go:225`

**Problem**: Direct indexing without bounds checking:
```go
return &(*(*resp.JSON200).Data)[0], nil
```

**Impact**: Panic if API returns empty array

**Solution**:
```go
data := *(*resp.JSON200).Data
if len(data) == 0 {
    return nil, errors.New("no data returned from API")
}
return &data[0], nil
```

---

## 📊 API Coverage Gaps

The library currently exposes only a **small subset** of available Hue API endpoints. Here are high-priority additions:

### Missing Core Resources

#### 1. **Zones** (Available but not wrapped)
```go
// Add to openhue.go
func (h *Home) GetZones(ctx context.Context) (map[string]ZoneGet, error)
func (h *Home) GetZoneById(ctx context.Context, zoneId string) (*ZoneGet, error)
func (h *Home) CreateZone(ctx context.Context, body CreateZoneJSONRequestBody) (*ZoneGet, error)
func (h *Home) UpdateZone(ctx context.Context, zoneId string, body UpdateZoneJSONRequestBody) error
func (h *Home) DeleteZone(ctx context.Context, zoneId string) error
```

#### 2. **Smart Scenes** (Available but not wrapped)
```go
func (h *Home) GetSmartScenes(ctx context.Context) (map[string]SmartSceneGet, error)
func (h *Home) CreateSmartScene(ctx context.Context, body CreateSmartSceneJSONRequestBody) (*SmartSceneGet, error)
func (h *Home) UpdateSmartScene(ctx context.Context, sceneId string, body UpdateSmartSceneJSONRequestBody) error
func (h *Home) DeleteSmartScene(ctx context.Context, sceneId string) error
```

#### 3. **Buttons/Switches** (Available but not wrapped)
```go
func (h *Home) GetButtons(ctx context.Context) (map[string]ButtonGet, error)
func (h *Home) GetButtonById(ctx context.Context, buttonId string) (*ButtonGet, error)
```

#### 4. **Motion Sensors** (Available but not wrapped)
```go
func (h *Home) GetMotionSensors(ctx context.Context) (map[string]MotionSensorGet, error)
func (h *Home) GetMotionSensorById(ctx context.Context, sensorId string) (*MotionSensorGet, error)
```

#### 5. **Temperature Sensors** (Available but not wrapped)
```go
func (h *Home) GetTemperatureSensors(ctx context.Context) (map[string]TemperatureGet, error)
func (h *Home) GetTemperatureSensorById(ctx context.Context, sensorId string) (*TemperatureGet, error)
```

#### 6. **Entertainment Areas** (Available but not wrapped)
```go
func (h *Home) GetEntertainmentConfigurations(ctx context.Context) (map[string]EntertainmentConfigurationGet, error)
func (h *Home) StartEntertainment(ctx context.Context, entertainmentId string) error
func (h *Home) StopEntertainment(ctx context.Context, entertainmentId string) error
```

#### 7. **Bridge Information** (Available but not wrapped)
```go
func (h *Home) GetBridge(ctx context.Context) (*BridgeGet, error)
func (h *Home) UpdateBridge(ctx context.Context, bridgeId string, body UpdateBridgeJSONRequestBody) error
```

#### 8. **Device Power & Battery** (Available but not wrapped)
```go
func (h *Home) GetDevicePowers(ctx context.Context) (map[string]DevicePowerGet, error)
func (h *Home) GetDevicePowerById(ctx context.Context, deviceId string) (*DevicePowerGet, error)
```

### Missing Scene Operations
```go
func (h *Home) GetSceneById(ctx context.Context, sceneId string) (*SceneGet, error)
func (h *Home) CreateScene(ctx context.Context, body CreateSceneJSONRequestBody) (*SceneGet, error)
func (h *Home) DeleteScene(ctx context.Context, sceneId string) error
func (h *Home) ActivateScene(ctx context.Context, sceneId string) error // Common use case!
```

### Missing Room Operations
```go
func (h *Home) GetRoomById(ctx context.Context, roomId string) (*RoomGet, error)
func (h *Home) CreateRoom(ctx context.Context, body CreateRoomJSONRequestBody) (*RoomGet, error)
func (h *Home) UpdateRoom(ctx context.Context, roomId string, body UpdateRoomJSONRequestBody) error
func (h *Home) DeleteRoom(ctx context.Context, roomId string) error
```

### Missing Light Operations
```go
func (h *Home) GetLightById(ctx context.Context, lightId string) (*LightGet, error)
```

---

## 🏗️ Code Quality Improvements

### 1. **Add Functional Options Pattern for Home**
Allow users to customize the HTTP client:

```go
type HomeOption func(*Home) error

func WithHTTPClient(client *http.Client) HomeOption {
    return func(h *Home) error {
        // Apply custom HTTP client
        return nil
    }
}

func WithTimeout(timeout time.Duration) HomeOption {
    return func(h *Home) error {
        // Apply timeout
        return nil
    }
}

func NewHome(bridgeIP, apiKey string, opts ...HomeOption) (*Home, error) {
    // Apply options
}
```

### 2. **Improve Error Types**
Create specific error types for better error handling:

```go
// error.go
type ErrNotFound struct {
    ResourceType string
    ResourceID   string
}

func (e *ErrNotFound) Error() string {
    return fmt.Sprintf("%s with ID %s not found", e.ResourceType, e.ResourceID)
}

type ErrInvalidResponse struct {
    StatusCode int
    Message    string
}

func (e *ErrInvalidResponse) Error() string {
    return fmt.Sprintf("invalid response (status %d): %s", e.StatusCode, e.Message)
}

// Allow users to check error types
if errors.Is(err, &ErrNotFound{}) {
    // Handle not found
}
```

### 3. **Add Proper Logging Interface**
```go
// Add to openhue.go
type Logger interface {
    Debug(msg string, keysAndValues ...interface{})
    Info(msg string, keysAndValues ...interface{})
    Error(msg string, keysAndValues ...interface{})
}

type Home struct {
    api    ClientWithResponsesInterface
    logger Logger
}

func WithLogger(logger Logger) HomeOption {
    return func(h *Home) error {
        h.logger = logger
        return nil
    }
}
```

### 4. **Consistent Return Types**
Some methods return pointers, others don't. Standardize:
- Collections: Return `map[string]Type` or `[]Type`
- Single items: Return `*Type` for consistency with nil checks

### 5. **Add Validation**
```go
func (h *Home) UpdateLight(ctx context.Context, lightId string, body LightPut) error {
    if lightId == "" {
        return errors.New("lightId cannot be empty")
    }
    // Continue...
}
```

---

## 🎯 Best Practices & Patterns

### 1. **Add Builder Pattern for Complex Types**
```go
// helpers.go
type LightUpdateBuilder struct {
    put LightPut
}

func NewLightUpdate() *LightUpdateBuilder {
    return &LightUpdateBuilder{}
}

func (b *LightUpdateBuilder) On(on bool) *LightUpdateBuilder {
    b.put.On = &On{On: &on}
    return b
}

func (b *LightUpdateBuilder) Brightness(brightness float64) *LightUpdateBuilder {
    b.put.Dimming = &Dimming{Brightness: &brightness}
    return b
}

func (b *LightUpdateBuilder) Build() LightPut {
    return b.put
}

// Usage:
update := openhue.NewLightUpdate().
    On(true).
    Brightness(80.0).
    Build()
```

### 2. **Add Convenience Methods**
```go
// helpers.go
func (h *Home) TurnOnLight(ctx context.Context, lightId string) error {
    on := true
    return h.UpdateLight(ctx, lightId, LightPut{
        On: &On{On: &on},
    })
}

func (h *Home) TurnOffLight(ctx context.Context, lightId string) error {
    off := false
    return h.UpdateLight(ctx, lightId, LightPut{
        On: &On{On: &off},
    })
}

func (h *Home) SetLightBrightness(ctx context.Context, lightId string, brightness float64) error {
    return h.UpdateLight(ctx, lightId, LightPut{
        Dimming: &Dimming{Brightness: &brightness},
    })
}

func (h *Home) TurnOnRoom(ctx context.Context, roomId string) error
func (h *Home) TurnOffRoom(ctx context.Context, roomId string) error
```

### 3. **Add Retry Logic**
```go
// Add to openhue.go
func (h *Home) withRetry(ctx context.Context, fn func() error) error {
    maxRetries := 3
    for i := 0; i < maxRetries; i++ {
        err := fn()
        if err == nil {
            return nil
        }
        if ctx.Err() != nil {
            return ctx.Err()
        }
        if i < maxRetries-1 {
            time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
        }
    }
    return errors.New("max retries exceeded")
}
```

### 4. **Add Resource Filtering**
```go
// helpers.go
type RoomFilter func(RoomGet) bool

func (h *Home) GetRoomsFiltered(ctx context.Context, filters ...RoomFilter) (map[string]RoomGet, error) {
    rooms, err := h.GetRooms(ctx)
    if err != nil {
        return nil, err
    }
    
    result := make(map[string]RoomGet)
    for id, room := range rooms {
        include := true
        for _, filter := range filters {
            if !filter(room) {
                include = false
                break
            }
        }
        if include {
            result[id] = room
        }
    }
    return result, nil
}

// Usage:
bedrooms := home.GetRoomsFiltered(ctx, func(r RoomGet) bool {
    return strings.Contains(*r.Metadata.Name, "Bedroom")
})
```

---

## 🧪 Testing Improvements

### 1. **Add Integration Tests**
```go
// openhue_integration_test.go
//go:build integration
// +build integration

func TestIntegration_RealBridge(t *testing.T) {
    // Tests that run against a real bridge
}
```

### 2. **Add More Unit Tests**
Currently only one test exists. Add comprehensive tests:
```go
func TestHome_GetDevices_Success(t *testing.T)
func TestHome_GetDevices_EmptyResponse(t *testing.T)
func TestHome_GetDevices_APIError(t *testing.T)
func TestHome_UpdateLight_Success(t *testing.T)
func TestHome_UpdateLight_NotFound(t *testing.T)
```

### 3. **Add Examples**
```go
// example_test.go
func ExampleHome_GetRooms() {
    home, _ := openhue.NewHome("192.168.1.2", "api-key")
    rooms, _ := home.GetRooms(context.Background())
    for _, room := range rooms {
        fmt.Println(*room.Metadata.Name)
    }
}
```

### 4. **Add Benchmarks**
```go
// openhue_bench_test.go
func BenchmarkHome_GetDevices(b *testing.B)
func BenchmarkHome_UpdateLight(b *testing.B)
```

---

## 📚 Documentation Improvements

### 1. **Expand Package Documentation**
`doc.go` is too brief. Add:
- Common use cases
- Error handling patterns
- Configuration examples
- Best practices

### 2. **Add Godoc Comments**
Many exported functions lack proper documentation:
```go
// GetDevices retrieves all devices connected to the Hue bridge.
// It returns a map where keys are device IDs and values are device details.
// 
// Example:
//   devices, err := home.GetDevices(ctx)
//   if err != nil {
//       log.Fatal(err)
//   }
//   for id, device := range devices {
//       fmt.Printf("Device %s: %s\n", id, *device.Metadata.Name)
//   }
func (h *Home) GetDevices(ctx context.Context) (map[string]DeviceGet, error)
```

### 3. **Add CONTRIBUTING.md**
Guide for contributors on:
- How to regenerate code
- Testing requirements
- PR process

---

## 🔧 Configuration Improvements

### 1. **Improve Config Loading**
`helpers.go` has issues:

```go
// Current: crashes on type assertion
c[\"bridge\"].(string)

// Better: handle errors gracefully
func LoadConf() (*Conf, error) {
    homedir, err := os.UserHomeDir()
    if err != nil {
        return nil, fmt.Errorf("unable to get home directory: %w", err)
    }

    yamlFile, err := os.ReadFile(filepath.Join(homedir, ".openhue", "config.yaml"))
    if err != nil {
        return nil, fmt.Errorf("unable to read config: %w", err)
    }

    var config struct {
        Bridge string `yaml:"bridge"`
        Key    string `yaml:"key"`
    }

    if err := yaml.Unmarshal(yamlFile, &config); err != nil {
        return nil, fmt.Errorf("unable to parse config: %w", err)
    }

    if config.Bridge == "" || config.Key == "" {
        return nil, errors.New("bridge and key must be set in config")
    }

    return &Conf{
        bridgeIP: config.Bridge,
        apiKey:   config.Key,
    }, nil
}
```

### 2. **Add Environment Variable Support**
```go
func LoadConfFromEnv() (*Conf, error) {
    bridge := os.Getenv("HUE_BRIDGE_IP")
    key := os.Getenv("HUE_API_KEY")
    
    if bridge == "" || key == "" {
        return nil, errors.New("HUE_BRIDGE_IP and HUE_API_KEY must be set")
    }
    
    return &Conf{bridgeIP: bridge, apiKey: key}, nil
}
```

### 3. **Add Config Precedence**
```go
func LoadConfAny() (*Conf, error) {
    // 1. Try environment variables
    if conf, err := LoadConfFromEnv(); err == nil {
        return conf, nil
    }
    
    // 2. Try config file
    if conf, err := LoadConf(); err == nil {
        return conf, nil
    }
    
    return nil, errors.New("no configuration found")
}
```

---

## 🚀 Additional Features

### 1. **Event Streaming (SSE)**
The Hue API supports Server-Sent Events for real-time updates:
```go
func (h *Home) StreamEvents(ctx context.Context) (<-chan Event, error)
```

### 2. **Bulk Operations**
```go
func (h *Home) UpdateMultipleLights(ctx context.Context, updates map[string]LightPut) error
func (h *Home) TurnOffAllLights(ctx context.Context) error
func (h *Home) TurnOffRoom(ctx context.Context, roomId string) error
```

### 3. **Scene Activation Helper**
```go
func (h *Home) ActivateScene(ctx context.Context, sceneId string) error {
    return h.UpdateScene(ctx, sceneId, ScenePut{
        Recall: &SceneRecall{
            Action: Ptr("active"),
        },
    })
}
```

### 4. **Color Helpers**
```go
func RGBToXY(r, g, b uint8) (x, y float64)
func XYToRGB(x, y float64) (r, g, b uint8)
func ColorTemperatureToMired(kelvin int) int
```

### 5. **Pointer Helpers**
```go
// helpers.go
func Ptr[T any](v T) *T {
    return &v
}

// Usage:
update := LightPut{
    On: &On{On: Ptr(true)},
    Dimming: &Dimming{Brightness: Ptr(80.0)},
}
```

---

## 📝 Priority Action Items

### High Priority (Security & Correctness)
1. ✅ Fix global HTTP transport mutation
2. ✅ Add context parameters to all methods
3. ✅ Fix unsafe array indexing
4. ✅ Add status code checking to all update methods

### Medium Priority (Usability)
5. Add Zones, Smart Scenes, Sensors wrappers
6. Add convenience methods (TurnOnLight, SetBrightness, etc.)
7. Improve error handling with specific error types
8. Add builder patterns for complex updates

### Low Priority (Nice to Have)
9. Add event streaming support
10. Add bulk operation helpers
11. Add color conversion utilities
12. Expand test coverage

---

## Conclusion

This library has a solid foundation but needs critical security fixes before production use. The auto-generation approach is excellent, but the manual wrappers need significant expansion to cover common use cases. Focus on:

1. **Security first**: Fix the global transport issue immediately
2. **API coverage**: Add wrappers for Zones, Sensors, and Smart Scenes
3. **Ergonomics**: Add convenience methods and builders
4. **Robustness**: Add proper context support and error handling

The library shows good potential and with these improvements would be excellent for production use.
