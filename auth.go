package openhue

import (
	"context"
	"errors"
	"fmt"
	"os"
)

// Authenticator defines a service that allows retrieving the Hue API Key
type Authenticator interface {
	// Authenticate performs a single authentication request to retrieve an API key.
	// It will return ("", true, err != nil) if the link button has not been pressed.
	Authenticate() (key string, press bool, err error)
}

type authenticatorImpl struct {
	client            *ClientWithResponses
	deviceType        string
	generateClientKey bool
}

type authOpt func(b *authenticatorImpl)

// NewAuthenticator creates an Authenticator targeting the Bridge at bridgeIP.
func NewAuthenticator(bridgeIP string, opts ...authOpt) (Authenticator, error) {
	client, err := newClient(bridgeIP, "", nil)
	if err != nil {
		return nil, err
	}

	authenticator := &authenticatorImpl{client: client, generateClientKey: true}

	for _, o := range opts {
		o(authenticator)
	}

	if len(authenticator.deviceType) == 0 {
		hostName, err := os.Hostname()
		if err != nil {
			return nil, err
		}
		authenticator.deviceType = hostName
	}

	return authenticator, nil
}

func (a *authenticatorImpl) Authenticate() (string, bool, error) {

	body := AuthenticateJSONRequestBody{
		Devicetype:        &a.deviceType,
		Generateclientkey: &a.generateClientKey,
	}

	resp, err := a.client.AuthenticateWithResponse(context.Background(), body)
	if err != nil {
		return "", false, err
	}

	if resp.JSON200 == nil {
		return "", false, fmt.Errorf("unable to reach the Bridge, verify that the IP is correct")
	}

	if len(*resp.JSON200) == 0 {
		return "", false, errors.New("empty authentication response from the Bridge")
	}

	auth := (*resp.JSON200)[0]
	if auth.Error != nil {
		desc := "link button not pressed"
		if auth.Error.Description != nil {
			desc = *auth.Error.Description
		}
		return "", true, errors.New(desc)
	}

	if auth.Success == nil || auth.Success.Username == nil || *auth.Success.Username == "" {
		return "", false, errors.New("authentication response did not contain an API key")
	}

	return *auth.Success.Username, false, nil
}

// WithDeviceType sets the device type reported to the Bridge. Defaults to the host name.
func WithDeviceType(deviceType string) authOpt {
	return func(b *authenticatorImpl) {
		b.deviceType = deviceType
	}
}

// WithGenerateClientKey controls whether the Bridge also generates a client key. Defaults to true.
func WithGenerateClientKey(generateClientKey bool) authOpt {
	return func(b *authenticatorImpl) {
		b.generateClientKey = generateClientKey
	}
}
