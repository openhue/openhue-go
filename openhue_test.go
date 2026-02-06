package openhue

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDevices_Forbidden(t *testing.T) {
	home, m := NewTestHome()

	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 403},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices(context.Background())
	assert.Error(t, err)
	assert.ErrorContains(t, err, "openhue api error (403)")
	assert.ErrorContains(t, err, "forbidden")
	assert.True(t, errors.Is(err, ErrForbidden))
}

func TestGetDevices_NotFound(t *testing.T) {
	home, m := NewTestHome()

	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 404},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices(context.Background())
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrNotFound))
}

func TestGetDevices_WithErrorDescription(t *testing.T) {
	home, m := NewTestHome()

	description := "resource not found: device xyz"
	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 404},
		JSON404: &NotFound{
			Errors: &[]Error{
				{Description: &description},
			},
		},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices(context.Background())
	assert.Error(t, err)
	assert.ErrorContains(t, err, "resource not found: device xyz")
	assert.True(t, errors.Is(err, ErrNotFound))
}

func TestGetDevices_Unauthorized(t *testing.T) {
	home, m := NewTestHome()

	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 401},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices(context.Background())
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrUnauthorized))
}

func TestGetDevices_TooManyRequests(t *testing.T) {
	home, m := NewTestHome()

	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 429},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices(context.Background())
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrTooManyRequests))
	assert.ErrorContains(t, err, "rate limited")
}

func TestApiError_Is(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		target     error
		want       bool
	}{
		{"401 is ErrUnauthorized", 401, ErrUnauthorized, true},
		{"403 is ErrForbidden", 403, ErrForbidden, true},
		{"404 is ErrNotFound", 404, ErrNotFound, true},
		{"409 is ErrConflict", 409, ErrConflict, true},
		{"429 is ErrTooManyRequests", 429, ErrTooManyRequests, true},
		{"500 is ErrInternalServerError", 500, ErrInternalServerError, true},
		{"503 is ErrServiceUnavailable", 503, ErrServiceUnavailable, true},
		{"403 is not ErrNotFound", 403, ErrNotFound, false},
		{"200 is not any sentinel", 200, ErrNotFound, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := &ApiError{StatusCode: tt.statusCode}
			got := errors.Is(err, tt.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetDeviceById_EmptyResponse(t *testing.T) {
	home, m := NewTestHome()

	emptyData := []DeviceGet{}
	resp := GetDeviceResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   *[]DeviceGet `json:"data,omitempty"`
			Errors *[]Error     `json:"errors,omitempty"`
		}{
			Data: &emptyData,
		},
	}
	m.On("GetDeviceWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDeviceById(context.Background(), "some-id")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrEmptyResponse))
}

func TestGetGroupedLightById_EmptyResponse(t *testing.T) {
	home, m := NewTestHome()

	emptyData := []GroupedLightGet{}
	resp := GetGroupedLightResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   *[]GroupedLightGet `json:"data,omitempty"`
			Errors *[]Error           `json:"errors,omitempty"`
		}{
			Data: &emptyData,
		},
	}
	m.On("GetGroupedLightWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetGroupedLightById(context.Background(), "some-id")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrEmptyResponse))
}

func TestGetBridgeHome_EmptyResponse(t *testing.T) {
	home, m := NewTestHome()

	emptyData := []BridgeHomeGet{}
	resp := GetBridgeHomesResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   *[]BridgeHomeGet `json:"data,omitempty"`
			Errors *[]Error         `json:"errors,omitempty"`
		}{
			Data: &emptyData,
		},
	}
	m.On("GetBridgeHomesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetBridgeHome(context.Background())
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrEmptyResponse))
}

func TestApiError_ErrorMessage(t *testing.T) {
	tests := []struct {
		name        string
		apiErr      *ApiError
		wantContain string
	}{
		{
			"with description",
			&ApiError{StatusCode: 404, Description: "light not found"},
			"light not found",
		},
		{
			"forbidden without description",
			&ApiError{StatusCode: 403},
			"wrong API key",
		},
		{
			"not found without description",
			&ApiError{StatusCode: 404},
			"resource not found",
		},
		{
			"unknown status code",
			&ApiError{StatusCode: 418, Status: "I'm a teapot"},
			"I'm a teapot",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, tt.apiErr.Error(), tt.wantContain)
		})
	}
}

// Tests for NewHome functional options

func TestNewHome_EmptyBridgeIP(t *testing.T) {
	_, err := NewHome("", "api-key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bridgeIP and apiKey must be set")
}

func TestNewHome_EmptyApiKey(t *testing.T) {
	_, err := NewHome("192.168.1.1", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "bridgeIP and apiKey must be set")
}

func TestWithCustomHTTPClient_NilClient(t *testing.T) {
	_, err := NewHome("192.168.1.1", "api-key", WithCustomHTTPClient(nil))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP client cannot be nil")
}

func TestWithRequestTimeout_NegativeTimeout(t *testing.T) {
	_, err := NewHome("192.168.1.1", "api-key", WithRequestTimeout(-1*time.Second))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "timeout cannot be negative")
}

func TestWithRequestTimeout_ZeroTimeout(t *testing.T) {
	// Zero timeout should be allowed (means no timeout)
	home, err := NewHome("192.168.1.1", "api-key", WithRequestTimeout(0))
	assert.NoError(t, err)
	assert.NotNil(t, home)
}

func TestWithCustomHTTPClient_ValidClient(t *testing.T) {
	customClient := &http.Client{}
	home, err := NewHome("192.168.1.1", "api-key", WithCustomHTTPClient(customClient))
	assert.NoError(t, err)
	assert.NotNil(t, home)
}

func TestWithRequestTimeout_ValidTimeout(t *testing.T) {
	home, err := NewHome("192.168.1.1", "api-key", WithRequestTimeout(60*time.Second))
	assert.NoError(t, err)
	assert.NotNil(t, home)
}

func TestNewHome_MultipleOptions(t *testing.T) {
	customClient := &http.Client{}
	home, err := NewHome("192.168.1.1", "api-key",
		WithCustomHTTPClient(customClient),
		WithRequestTimeout(30*time.Second), // This will be ignored since custom client is set
	)
	assert.NoError(t, err)
	assert.NotNil(t, home)
}

func TestNewHome_DefaultTimeout(t *testing.T) {
	// Test that NewHome works without any options
	home, err := NewHome("192.168.1.1", "api-key")
	assert.NoError(t, err)
	assert.NotNil(t, home)
}
