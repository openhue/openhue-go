package openhue

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
)

type Home struct {
	api ClientWithResponsesInterface
}

// NewHome creates a new Home context that is able to manage your different Philips Hue devices.
func NewHome(bridgeIP, apiKey string) (*Home, error) {

	if bridgeIP == "" || apiKey == "" {
		return nil, errors.New("illegal arguments, bridgeIP and apiKey must be set")
	}

	client, err := newClient(bridgeIP, apiKey)
	if err != nil {
		return nil, err
	}

	return &Home{
		api: client,
	}, nil
}

//--------------------------------------------------------------------------------------------------------------------//
// BRIDGE HOME
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetBridgeHome(ctx context.Context) (*BridgeHomeGet, error) {

	resp, err := h.api.GetBridgeHomesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}
	if len(data) != 1 {
		return nil, errors.New("more than 1 home attached to the bridge is not supported yet")
	}

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// RESOURCE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetResources(ctx context.Context) (map[string]ResourceGet, error) {
	resp, err := h.api.GetResourcesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	resources := make(map[string]ResourceGet)

	for _, resource := range data {
		resources[*resource.Id] = resource
	}

	return resources, nil
}

//--------------------------------------------------------------------------------------------------------------------//
// DEVICE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetDevices(ctx context.Context) (map[string]DeviceGet, error) {
	resp, err := h.api.GetDevicesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	devices := make(map[string]DeviceGet)

	for _, device := range data {
		devices[*device.Id] = device
	}

	return devices, nil
}

func (h *Home) GetDeviceById(ctx context.Context, deviceId string) (*DeviceGet, error) {
	resp, err := h.api.GetDeviceWithResponse(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// ROOM
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetRooms(ctx context.Context) (map[string]RoomGet, error) {
	resp, err := h.api.GetRoomsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	rooms := make(map[string]RoomGet)

	for _, room := range data {
		rooms[*room.Id] = room
	}

	return rooms, nil
}

func (h *Home) GetRoomById(ctx context.Context, roomId string) (*RoomGet, error) {
	resp, err := h.api.GetRoomWithResponse(ctx, roomId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) CreateRoom(ctx context.Context, body RoomPut) (*ResourceIdentifier, error) {
	resp, err := h.api.CreateRoomWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) UpdateRoom(ctx context.Context, roomId string, body RoomPut) error {
	resp, err := h.api.UpdateRoomWithResponse(ctx, roomId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

func (h *Home) DeleteRoom(ctx context.Context, roomId string) error {
	resp, err := h.api.DeleteRoomWithResponse(ctx, roomId)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

func (r *RoomGet) GetServices() map[string]ResourceIdentifierRtype {
	services := make(map[string]ResourceIdentifierRtype)
	for _, s := range *r.Services {
		services[*s.Rid] = *s.Rtype
	}
	return services
}

//--------------------------------------------------------------------------------------------------------------------//
// ZONE
// Note: Zones use RoomGet/RoomPut types in the API
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetZones(ctx context.Context) (map[string]RoomGet, error) {
	resp, err := h.api.GetZonesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	zones := make(map[string]RoomGet)

	for _, zone := range data {
		zones[*zone.Id] = zone
	}

	return zones, nil
}

func (h *Home) GetZoneById(ctx context.Context, zoneId string) (*RoomGet, error) {
	resp, err := h.api.GetZoneWithResponse(ctx, zoneId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) CreateZone(ctx context.Context, body RoomPut) (*ResourceIdentifier, error) {
	resp, err := h.api.CreateZoneWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) UpdateZone(ctx context.Context, zoneId string, body RoomPut) error {
	resp, err := h.api.UpdateZoneWithResponse(ctx, zoneId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

func (h *Home) DeleteZone(ctx context.Context, zoneId string) error {
	resp, err := h.api.DeleteZoneWithResponse(ctx, zoneId)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetLights(ctx context.Context) (map[string]LightGet, error) {
	resp, err := h.api.GetLightsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	lights := make(map[string]LightGet)

	for _, light := range data {
		lights[*light.Id] = light
	}

	return lights, nil
}

func (h *Home) GetLightById(ctx context.Context, lightId string) (*LightGet, error) {
	resp, err := h.api.GetLightWithResponse(ctx, lightId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (l *LightGet) IsOn() bool {
	return *l.On.On
}

func (l *LightGet) Toggle() *On {
	v := !(*l.On.On)
	return &On{On: &v}
}

func (h *Home) UpdateLight(ctx context.Context, lightId string, body LightPut) error {
	resp, err := h.api.UpdateLightWithResponse(ctx, lightId, body)
	if err != nil {
		return err
	}
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}
	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// GROUPED LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetGroupedLights(ctx context.Context) (map[string]GroupedLightGet, error) {
	resp, err := h.api.GetGroupedLightsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	lights := make(map[string]GroupedLightGet)

	for _, light := range data {
		lights[*light.Id] = light
	}

	return lights, nil
}

func (h *Home) GetGroupedLightById(ctx context.Context, groupedLightId string) (*GroupedLightGet, error) {
	resp, err := h.api.GetGroupedLightWithResponse(ctx, groupedLightId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (l *GroupedLightGet) IsOn() bool {
	return *l.On.On
}

func (l *GroupedLightGet) Toggle() *On {
	v := !(*l.On.On)
	return &On{On: &v}
}

func (h *Home) UpdateGroupedLight(ctx context.Context, lightId string, body GroupedLightPut) error {
	resp, err := h.api.UpdateGroupedLightWithResponse(ctx, lightId, body)
	if err != nil {
		return err
	}
	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}
	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// SCENE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetScenes(ctx context.Context) (map[string]SceneGet, error) {
	resp, err := h.api.GetScenesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	scenes := make(map[string]SceneGet)

	for _, scene := range data {
		scenes[*scene.Id] = scene
	}

	return scenes, nil
}

func (h *Home) GetSceneById(ctx context.Context, sceneId string) (*SceneGet, error) {
	resp, err := h.api.GetSceneWithResponse(ctx, sceneId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) CreateScene(ctx context.Context, body ScenePost) (*ResourceIdentifier, error) {
	resp, err := h.api.CreateSceneWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) UpdateScene(ctx context.Context, sceneId string, body ScenePut) error {
	resp, err := h.api.UpdateSceneWithResponse(ctx, sceneId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

func (h *Home) DeleteScene(ctx context.Context, sceneId string) error {
	resp, err := h.api.DeleteSceneWithResponse(ctx, sceneId)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

// ActivateScene activates a scene by its ID.
func (h *Home) ActivateScene(ctx context.Context, sceneId string) error {
	action := SceneRecallActionActive
	return h.UpdateScene(ctx, sceneId, ScenePut{
		Recall: &SceneRecall{
			Action: &action,
		},
	})
}

//--------------------------------------------------------------------------------------------------------------------//
// SMART SCENE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetSmartScenes(ctx context.Context) (map[string]SmartSceneGet, error) {
	resp, err := h.api.GetSmartScenesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	scenes := make(map[string]SmartSceneGet)

	for _, scene := range data {
		scenes[*scene.Id] = scene
	}

	return scenes, nil
}

func (h *Home) GetSmartSceneById(ctx context.Context, smartSceneId string) (*SmartSceneGet, error) {
	resp, err := h.api.GetSmartSceneWithResponse(ctx, smartSceneId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) CreateSmartScene(ctx context.Context, body SmartScenePost) (*ResourceIdentifier, error) {
	resp, err := h.api.CreateSmartSceneWithResponse(ctx, body)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

func (h *Home) UpdateSmartScene(ctx context.Context, smartSceneId string, body SmartScenePut) error {
	resp, err := h.api.UpdateSmartSceneWithResponse(ctx, smartSceneId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

func (h *Home) DeleteSmartScene(ctx context.Context, smartSceneId string) error {
	resp, err := h.api.DeleteSmartSceneWithResponse(ctx, smartSceneId)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// BUTTON
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetButtons(ctx context.Context) (map[string]ButtonGet, error) {
	resp, err := h.api.GetButtonsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	buttons := make(map[string]ButtonGet)

	for _, button := range data {
		buttons[*button.Id] = button
	}

	return buttons, nil
}

func (h *Home) GetButtonById(ctx context.Context, buttonId string) (*ButtonGet, error) {
	resp, err := h.api.GetButtonWithResponse(ctx, buttonId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// MOTION SENSOR
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetMotionSensors(ctx context.Context) (map[string]MotionGet, error) {
	resp, err := h.api.GetMotionSensorsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	sensors := make(map[string]MotionGet)

	for _, sensor := range data {
		sensors[*sensor.Id] = sensor
	}

	return sensors, nil
}

func (h *Home) GetMotionSensorById(ctx context.Context, motionSensorId string) (*MotionGet, error) {
	resp, err := h.api.GetMotionSensorWithResponse(ctx, motionSensorId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// TEMPERATURE SENSOR
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetTemperatureSensors(ctx context.Context) (map[string]TemperatureGet, error) {
	resp, err := h.api.GetTemperaturesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	sensors := make(map[string]TemperatureGet)

	for _, sensor := range data {
		sensors[*sensor.Id] = sensor
	}

	return sensors, nil
}

func (h *Home) GetTemperatureSensorById(ctx context.Context, temperatureId string) (*TemperatureGet, error) {
	resp, err := h.api.GetTemperatureWithResponse(ctx, temperatureId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// ENTERTAINMENT CONFIGURATION
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetEntertainmentConfigurations(ctx context.Context) (map[string]EntertainmentConfigurationGet, error) {
	resp, err := h.api.GetEntertainmentConfigurationsWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	configs := make(map[string]EntertainmentConfigurationGet)

	for _, config := range data {
		configs[*config.Id] = config
	}

	return configs, nil
}

func (h *Home) GetEntertainmentConfigurationById(ctx context.Context, entertainmentConfigurationId string) (*EntertainmentConfigurationGet, error) {
	resp, err := h.api.GetEntertainmentConfigurationWithResponse(ctx, entertainmentConfigurationId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

// StartEntertainment starts entertainment mode for the given configuration.
func (h *Home) StartEntertainment(ctx context.Context, entertainmentConfigurationId string) error {
	action := EntertainmentConfigurationPutActionStart
	return h.UpdateEntertainmentConfiguration(ctx, entertainmentConfigurationId, EntertainmentConfigurationPut{
		Action: &action,
	})
}

// StopEntertainment stops entertainment mode for the given configuration.
func (h *Home) StopEntertainment(ctx context.Context, entertainmentConfigurationId string) error {
	action := EntertainmentConfigurationPutActionStop
	return h.UpdateEntertainmentConfiguration(ctx, entertainmentConfigurationId, EntertainmentConfigurationPut{
		Action: &action,
	})
}

func (h *Home) UpdateEntertainmentConfiguration(ctx context.Context, entertainmentConfigurationId string, body EntertainmentConfigurationPut) error {
	resp, err := h.api.UpdateEntertainmentConfigurationWithResponse(ctx, entertainmentConfigurationId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// BRIDGE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetBridges(ctx context.Context) (map[string]BridgeGet, error) {
	resp, err := h.api.GetBridgesWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	bridges := make(map[string]BridgeGet)

	for _, bridge := range data {
		bridges[*bridge.Id] = bridge
	}

	return bridges, nil
}

func (h *Home) GetBridge(ctx context.Context) (*BridgeGet, error) {
	bridges, err := h.GetBridges(ctx)
	if err != nil {
		return nil, err
	}

	if len(bridges) == 0 {
		return nil, ErrEmptyResponse
	}

	// Return the first (and typically only) bridge
	for _, bridge := range bridges {
		return &bridge, nil
	}

	return nil, ErrEmptyResponse
}

func (h *Home) UpdateBridge(ctx context.Context, bridgeId string, body BridgePut) error {
	resp, err := h.api.UpdateBridgeWithResponse(ctx, bridgeId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// DEVICE POWER
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetDevicePowers(ctx context.Context) (map[string]DevicePowerGet, error) {
	resp, err := h.api.GetDevicePowersWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	powers := make(map[string]DevicePowerGet)

	for _, power := range data {
		powers[*power.Id] = power
	}

	return powers, nil
}

func (h *Home) GetDevicePowerById(ctx context.Context, devicePowerId string) (*DevicePowerGet, error) {
	resp, err := h.api.GetDevicePowerWithResponse(ctx, devicePowerId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data
	if len(data) == 0 {
		return nil, ErrEmptyResponse
	}

	return &data[0], nil
}

//
// Internal
//

// newClient creates a new ClientWithResponses for a given Bridge IP and API key.
// This function configures TLS to trust the Philips Hue Bridge root CA certificates.
func newClient(bridgeIP, apiKey string) (*ClientWithResponses, error) {

	// Create a certificate pool with the Hue Bridge root CA certificates
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM([]byte(HueBridgeRootCAs)) {
		return nil, errors.New("failed to parse Hue Bridge root CA certificates")
	}

	// Clone the default transport to preserve defaults like ProxyFromEnvironment,
	// timeouts, and HTTP/2 support. Only override TLS configuration.
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{
		RootCAs: certPool,
		// linter ignore:go/disabled-certificate-check
		InsecureSkipVerify: true,
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			// Manually verify the certificate chain against the Hue root CAs
			if len(rawCerts) == 0 {
				return errors.New("no certificates presented")
			}

			// Parse the leaf certificate
			cert, err := x509.ParseCertificate(rawCerts[0])
			if err != nil {
				return err
			}

			// Create intermediate pool from remaining certificates
			intermediates := x509.NewCertPool()
			for _, certBytes := range rawCerts[1:] {
				intermediate, err := x509.ParseCertificate(certBytes)
				if err != nil {
					return err
				}
				intermediates.AddCert(intermediate)
			}

			// Verify the certificate chain
			opts := x509.VerifyOptions{
				Roots:         certPool,
				Intermediates: intermediates,
				KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			}
			_, err = cert.Verify(opts)
			return err
		},
	}

	// Create HTTP client with custom transport
	httpClient := &http.Client{
		Transport: customTransport,
	}

	var authFn RequestEditorFn
	if len(apiKey) > 0 {
		authFn = func(ctx context.Context, req *http.Request) error {
			req.Header.Set("hue-application-key", apiKey)
			return nil
		}
	} else {
		authFn = func(ctx context.Context, req *http.Request) error {
			return nil
		}
	}

	return NewClientWithResponses(
		"https://"+bridgeIP,
		WithHTTPClient(httpClient),
		WithRequestEditorFn(authFn),
	)
}
