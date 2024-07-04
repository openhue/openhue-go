package openhue

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestNewHome(t *testing.T) {

	home, m := NewTestHome()

	resp := GetDevicesResponse{
		HTTPResponse: &http.Response{StatusCode: 403},
	}
	m.On("GetDevicesWithResponse", mock.Anything, mock.Anything).Return(&resp, nil)

	_, err := home.GetDevices()
	assert.Error(t, err)
	assert.ErrorContains(t, err, "openhue api error: wrong API key")
}
