package openhue

import (
	"context"
	"crypto/tls"
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

func (h *Home) GetBridgeHome() (*BridgeHomeGet, error) {

	resp, err := h.api.GetBridgeHomesWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	if len(*(*resp.JSON200).Data) != 1 {
		return nil, errors.New("more than 1 home attached to the bridge is not supported yet")
	}

	return &(*(*resp.JSON200).Data)[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// RESOURCE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetResources() (map[string]ResourceGet, error) {
	resp, err := h.api.GetResourcesWithResponse(context.Background())
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

func (h *Home) GetDevices() (map[string]DeviceGet, error) {
	resp, err := h.api.GetDevicesWithResponse(context.Background())
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

func (h *Home) GetDeviceById(deviceId string) (*DeviceGet, error) {
	resp, err := h.api.GetDeviceWithResponse(context.Background(), deviceId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data

	return &data[0], nil
}

//--------------------------------------------------------------------------------------------------------------------//
// ROOM
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetRooms() (map[string]RoomGet, error) {
	resp, err := h.api.GetRoomsWithResponse(context.Background())
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

func (r *RoomGet) GetServices() map[string]ResourceIdentifierRtype {
	services := make(map[string]ResourceIdentifierRtype)
	for _, s := range *r.Services {
		services[*s.Rid] = *s.Rtype
	}
	return services
}

//--------------------------------------------------------------------------------------------------------------------//
// LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetLights() (map[string]LightGet, error) {
	resp, err := h.api.GetLightsWithResponse(context.Background())
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

func (l *LightGet) IsOn() bool {
	return *l.On.On
}

func (l *LightGet) Toggle() *On {
	v := !(*l.On.On)
	return &On{On: &v}
}

func (h *Home) UpdateLight(lightId string, body LightPut) error {
	_, err := h.api.UpdateLightWithResponse(context.Background(), lightId, body)
	if err != nil {
		return err
	}
	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// GROUPED LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetGroupedLights() (map[string]GroupedLightGet, error) {
	resp, err := h.api.GetGroupedLightsWithResponse(context.Background())
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

func (h *Home) GetGroupedLightById(groupedLightId string) (*GroupedLightGet, error) {
	resp, err := h.api.GetGroupedLightWithResponse(context.Background(), groupedLightId)
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	data := *(*resp.JSON200).Data

	return &data[0], nil
}

func (l *GroupedLightGet) IsOn() bool {
	return *l.On.On
}

func (l *GroupedLightGet) Toggle() *On {
	v := !(*l.On.On)
	return &On{On: &v}
}

func (h *Home) UpdateGroupedLight(lightId string, body GroupedLightPut) error {
	_, err := h.api.UpdateGroupedLightWithResponse(context.Background(), lightId, body)
	if err != nil {
		return err
	}
	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// SCENE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetScenes() (map[string]SceneGet, error) {
	resp, err := h.api.GetScenesWithResponse(context.Background())
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

func (h *Home) UpdateScene(sceneId string, body ScenePut) error {
	resp, err := h.api.UpdateSceneWithResponse(context.Background(), sceneId, body)
	if err != nil {
		return err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return newApiError(resp)
	}

	return nil
}

//
// Internal
//

// newClient creates a new ClientWithResponses for a given Bridge IP and API key.
// This function will also skip SSL verification, as the Philips HUE Bridge exposes a self-signed certificate.
func newClient(bridgeIP, apiKey string) (*ClientWithResponses, error) {

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

	// skip SSL Verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	return NewClientWithResponses("https://"+bridgeIP, WithRequestEditorFn(authFn))
}
