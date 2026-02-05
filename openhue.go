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

	if len(*(*resp.JSON200).Data) != 1 {
		return nil, errors.New("more than 1 home attached to the bridge is not supported yet")
	}

	return &(*(*resp.JSON200).Data)[0], nil
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
