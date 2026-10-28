package openhue

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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

	// Decode from JSON so the test holds whatever Go shape oapi-codegen generates for Errors/Description.
	notFound := &NotFound{}
	require.NoError(t, json.Unmarshal([]byte(`{"errors":[{"description":"resource not found: device xyz"}]}`), notFound))
	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 404},
		JSON404:      notFound,
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

	resp := GetDeviceResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	// Decode from JSON so the test holds whatever Go shape oapi-codegen generates for Errors.
	require.NoError(t, json.Unmarshal([]byte(`{"data":[]}`), &resp.JSON200))
	m.On("GetDeviceWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDeviceById(context.Background(), "some-id")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrEmptyResponse))
}

func TestGetGroupedLightById_EmptyResponse(t *testing.T) {
	home, m := NewTestHome()

	resp := GetGroupedLightResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	// Decode from JSON so the test holds whatever Go shape oapi-codegen generates for Errors.
	require.NoError(t, json.Unmarshal([]byte(`{"data":[]}`), &resp.JSON200))
	m.On("GetGroupedLightWithResponse", mock.Anything, mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetGroupedLightById(context.Background(), "some-id")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, ErrEmptyResponse))
}

func TestGetBridgeHome_EmptyResponse(t *testing.T) {
	home, m := NewTestHome()

	resp := GetBridgeHomesResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	// Decode from JSON so the test holds whatever Go shape oapi-codegen generates for Errors.
	require.NoError(t, json.Unmarshal([]byte(`{"data":[]}`), &resp.JSON200))
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

func TestLightGet_IsOnAndToggle(t *testing.T) {
	var l LightGet
	assert.False(t, l.IsOn())
	assert.True(t, l.Toggle().On)

	l.On = On{On: true}
	assert.True(t, l.IsOn())
	assert.False(t, l.Toggle().On)
}

func TestGroupedLightGet_IsOnAndToggle_NilSafe(t *testing.T) {
	var l GroupedLightGet
	assert.False(t, l.IsOn())
	assert.True(t, l.Toggle().On)

	l.On = &On{On: true}
	assert.True(t, l.IsOn())
	assert.False(t, l.Toggle().On)
}

func TestGetDevices_NilBody(t *testing.T) {
	home, mockClient := NewTestHome()
	mockClient.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(
		&GetDevicesResponse{HTTPResponse: &http.Response{StatusCode: http.StatusOK}}, nil)

	_, err := home.GetDevices(context.Background())
	assert.ErrorIs(t, err, ErrEmptyResponse)
}

func TestGetRelativeRotaries(t *testing.T) {
	home, m := NewTestHome()

	resp := GetRelativeRotariesResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   []RelativeRotaryGet `json:"data"`
			Errors []Error             `json:"errors"`
		}{Data: []RelativeRotaryGet{{Id: "r1"}, {Id: "r2"}}},
	}
	m.On("GetRelativeRotariesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	got, err := home.GetRelativeRotaries(context.Background())
	require.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "r1", got["r1"].Id)
}

func TestGetRelativeRotaryById(t *testing.T) {
	home, m := NewTestHome()

	resp := GetRelativeRotaryResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   []RelativeRotaryGet `json:"data"`
			Errors []Error             `json:"errors"`
		}{Data: []RelativeRotaryGet{{Id: "r1"}}},
	}
	m.On("GetRelativeRotaryWithResponse", mock.Anything, "r1", mock.Anything).Return(&resp, nil)

	got, err := home.GetRelativeRotaryById(context.Background(), "r1")
	require.NoError(t, err)
	assert.Equal(t, "r1", got.Id)
}

func TestGetRelativeRotaryById_NotFound(t *testing.T) {
	home, m := NewTestHome()

	resp := GetRelativeRotaryResponse{HTTPResponse: &http.Response{StatusCode: 404}}
	m.On("GetRelativeRotaryWithResponse", mock.Anything, "x", mock.Anything).Return(&resp, nil)

	_, err := home.GetRelativeRotaryById(context.Background(), "x")
	assert.True(t, errors.Is(err, ErrNotFound))
}

func TestUpdateButton(t *testing.T) {
	home, m := NewTestHome()

	interval := 800
	body := ButtonPut{Button: &struct {
		RepeatInterval *int `json:"repeat_interval,omitempty"`
	}{RepeatInterval: &interval}}
	resp := UpdateButtonResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	m.On("UpdateButtonWithResponse", mock.Anything, "b1", body, mock.Anything).Return(&resp, nil)

	require.NoError(t, home.UpdateButton(context.Background(), "b1", body))
	m.AssertExpectations(t)
}

func TestUpdateButton_Error(t *testing.T) {
	home, m := NewTestHome()

	resp := UpdateButtonResponse{HTTPResponse: &http.Response{StatusCode: 403}}
	m.On("UpdateButtonWithResponse", mock.Anything, "b1", mock.Anything, mock.Anything).Return(&resp, nil)

	err := home.UpdateButton(context.Background(), "b1", ButtonPut{})
	assert.True(t, errors.Is(err, ErrForbidden))
}

func TestUpdateRelativeRotary(t *testing.T) {
	home, m := NewTestHome()

	resp := UpdateRelativeRotaryResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	m.On("UpdateRelativeRotaryWithResponse", mock.Anything, "r1", mock.Anything, mock.Anything).Return(&resp, nil)

	require.NoError(t, home.UpdateRelativeRotary(context.Background(), "r1", RelativeRotaryPut{}))
}

func TestGetBehaviorInstances(t *testing.T) {
	home, m := NewTestHome()

	resp := GetBehaviorInstancesResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   []BehaviorInstanceGet `json:"data"`
			Errors []Error               `json:"errors"`
		}{Data: []BehaviorInstanceGet{{Id: "bi1"}, {Id: "bi2"}}},
	}
	m.On("GetBehaviorInstancesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	got, err := home.GetBehaviorInstances(context.Background())
	require.NoError(t, err)
	assert.Len(t, got, 2)
	assert.Equal(t, "bi1", got["bi1"].Id)
}

func TestGetBehaviorInstanceById(t *testing.T) {
	home, m := NewTestHome()

	resp := GetBehaviorInstanceResponse{
		HTTPResponse: &http.Response{StatusCode: 200},
		JSON200: &struct {
			Data   []BehaviorInstanceGet `json:"data"`
			Errors []Error               `json:"errors"`
		}{Data: []BehaviorInstanceGet{{Id: "bi1"}}},
	}
	m.On("GetBehaviorInstanceWithResponse", mock.Anything, "bi1", mock.Anything).Return(&resp, nil)

	got, err := home.GetBehaviorInstanceById(context.Background(), "bi1")
	require.NoError(t, err)
	assert.Equal(t, "bi1", got.Id)
}

func TestGetBehaviorInstanceById_NotFound(t *testing.T) {
	home, m := NewTestHome()

	resp := GetBehaviorInstanceResponse{HTTPResponse: &http.Response{StatusCode: 404}}
	m.On("GetBehaviorInstanceWithResponse", mock.Anything, "x", mock.Anything).Return(&resp, nil)

	_, err := home.GetBehaviorInstanceById(context.Background(), "x")
	assert.True(t, errors.Is(err, ErrNotFound))
}

func TestUpdateBehaviorInstance(t *testing.T) {
	home, m := NewTestHome()

	cfg := map[string]interface{}{"buttons": map[string]interface{}{}}
	body := BehaviorInstancePut{Configuration: &cfg}
	resp := UpdateBehaviorInstanceResponse{HTTPResponse: &http.Response{StatusCode: 200}}
	m.On("UpdateBehaviorInstanceWithResponse", mock.Anything, "bi1", body, mock.Anything).Return(&resp, nil)

	require.NoError(t, home.UpdateBehaviorInstance(context.Background(), "bi1", body))
	m.AssertExpectations(t)
}

func TestUpdateBehaviorInstance_Error(t *testing.T) {
	home, m := NewTestHome()

	resp := UpdateBehaviorInstanceResponse{HTTPResponse: &http.Response{StatusCode: 403}}
	m.On("UpdateBehaviorInstanceWithResponse", mock.Anything, "bi1", mock.Anything, mock.Anything).Return(&resp, nil)

	err := home.UpdateBehaviorInstance(context.Background(), "bi1", BehaviorInstancePut{})
	assert.True(t, errors.Is(err, ErrForbidden))
}
