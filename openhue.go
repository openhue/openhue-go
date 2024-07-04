package openhue

import (
	"context"
	"crypto/tls"
	"errors"
	sp "github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	"log/slog"
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
		slog.Error("Error creating new Home client", err)
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
// DEVICE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetDevices() ([]DeviceGet, error) {
	resp, err := h.api.GetDevicesWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	return *(*resp.JSON200).Data, nil
}

//--------------------------------------------------------------------------------------------------------------------//
// ROOM
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetRooms() ([]RoomGet, error) {
	resp, err := h.api.GetRoomsWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	return *(*resp.JSON200).Data, nil
}

//--------------------------------------------------------------------------------------------------------------------//
// LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetLights() ([]LightGet, error) {
	resp, err := h.api.GetLightsWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	return *(*resp.JSON200).Data, nil
}

func (l *LightGet) IsOn() bool {
	return *l.On.On
}

func (h *Home) SetLight(lightId string, body UpdateLightJSONRequestBody) error {
	_, err := h.api.UpdateLightWithResponse(context.Background(), lightId, body)
	if err != nil {
		return err
	}
	return nil
}

//--------------------------------------------------------------------------------------------------------------------//
// GROUPED LIGHT
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetGroupedLights() ([]GroupedLightGet, error) {
	resp, err := h.api.GetGroupedLightsWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	return *(*resp.JSON200).Data, nil
}

//--------------------------------------------------------------------------------------------------------------------//
// SCENE
//--------------------------------------------------------------------------------------------------------------------//

func (h *Home) GetScenes() ([]SceneGet, error) {
	resp, err := h.api.GetScenesWithResponse(context.Background())
	if err != nil {
		return nil, err
	}

	if resp.HTTPResponse.StatusCode != http.StatusOK {
		return nil, newApiError(resp)
	}

	return *(*resp.JSON200).Data, nil
}

// NewOpenHueClient Creates a new NewClientWithResponses for a given server and hueApplicationKey.
// This function will also skip SSL verification, as the Philips HUE Bridge exposes a self-signed certificate.
func newClient(bridgeIP, apiKey string) (*ClientWithResponses, error) {
	p, err := sp.NewSecurityProviderApiKey("header", "hue-application-Key", apiKey)
	if err != nil {
		return nil, err
	}
	// skip SSL Verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client, err := NewClientWithResponses("https://"+bridgeIP, WithRequestEditorFn(p.Intercept))
	if err != nil {
		return nil, err
	}

	return client, nil
}
