package openhue

import (
	"context"
	"github.com/stretchr/testify/mock"
	"io"
)

type ClientWithResponsesMock struct {
	mock.Mock
}

func NewTestHome() (*Home, *ClientWithResponsesMock) {
	apiClientMock := new(ClientWithResponsesMock)
	return &Home{
		api: apiClientMock,
	}, apiClientMock
}

func (c *ClientWithResponsesMock) AuthenticateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error) {
	args := c.Called(ctx, contentType, body, reqEditors)
	return args.Get(0).(*AuthenticateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) AuthenticateWithResponse(ctx context.Context, body AuthenticateJSONRequestBody, reqEditors ...RequestEditorFn) (*AuthenticateResponse, error) {
	args := c.Called(ctx, body, reqEditors)
	return args.Get(0).(*AuthenticateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetResourcesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetResourcesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetResourcesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBehaviorInstancesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBehaviorInstancesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetBehaviorInstancesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBehaviorInstanceWithResponse(ctx context.Context, behaviorInstanceId string, reqEditors ...RequestEditorFn) (*GetBehaviorInstanceResponse, error) {
	args := c.Called(ctx, behaviorInstanceId, reqEditors)
	return args.Get(0).(*GetBehaviorInstanceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBehaviorInstanceWithBodyWithResponse(ctx context.Context, behaviorInstanceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateBehaviorInstanceResponse, error) {
	args := c.Called(ctx, behaviorInstanceId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateBehaviorInstanceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBehaviorInstanceWithResponse(ctx context.Context, behaviorInstanceId string, body UpdateBehaviorInstanceJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateBehaviorInstanceResponse, error) {
	args := c.Called(ctx, behaviorInstanceId, body, reqEditors)
	return args.Get(0).(*UpdateBehaviorInstanceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBehaviorScriptsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBehaviorScriptsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetBehaviorScriptsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBehaviorScriptWithResponse(ctx context.Context, behaviorScriptId string, reqEditors ...RequestEditorFn) (*GetBehaviorScriptResponse, error) {
	args := c.Called(ctx, behaviorScriptId, reqEditors)
	return args.Get(0).(*GetBehaviorScriptResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBehaviorScriptWithBodyWithResponse(ctx context.Context, behaviorScriptId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateBehaviorScriptResponse, error) {
	args := c.Called(ctx, behaviorScriptId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateBehaviorScriptResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBehaviorScriptWithResponse(ctx context.Context, behaviorScriptId string, body UpdateBehaviorScriptJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateBehaviorScriptResponse, error) {
	args := c.Called(ctx, behaviorScriptId, body, reqEditors)
	return args.Get(0).(*UpdateBehaviorScriptResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBellButtonsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBellButtonsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetBellButtonsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBellButtonWithResponse(ctx context.Context, bellButtonId string, reqEditors ...RequestEditorFn) (*GetBellButtonResponse, error) {
	args := c.Called(ctx, bellButtonId, reqEditors)
	return args.Get(0).(*GetBellButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBellButtonWithBodyWithResponse(ctx context.Context, bellButtonId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateBellButtonResponse, error) {
	args := c.Called(ctx, bellButtonId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateBellButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBellButtonWithResponse(ctx context.Context, bellButtonId string, body UpdateBellButtonJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateBellButtonResponse, error) {
	args := c.Called(ctx, bellButtonId, body, reqEditors)
	return args.Get(0).(*UpdateBellButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBridgesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBridgesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetBridgesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBridgeWithResponse(ctx context.Context, bridgeId string, reqEditors ...RequestEditorFn) (*GetBridgeResponse, error) {
	args := c.Called(ctx, bridgeId, reqEditors)
	return args.Get(0).(*GetBridgeResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBridgeWithBodyWithResponse(ctx context.Context, bridgeId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateBridgeResponse, error) {
	args := c.Called(ctx, bridgeId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateBridgeResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateBridgeWithResponse(ctx context.Context, bridgeId string, body UpdateBridgeJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateBridgeResponse, error) {
	args := c.Called(ctx, bridgeId, body, reqEditors)
	return args.Get(0).(*UpdateBridgeResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBridgeHomesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetBridgeHomesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetBridgeHomesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetBridgeHomeWithResponse(ctx context.Context, bridgeHomeId string, reqEditors ...RequestEditorFn) (*GetBridgeHomeResponse, error) {
	args := c.Called(ctx, bridgeHomeId, reqEditors)
	return args.Get(0).(*GetBridgeHomeResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetButtonsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetButtonsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetButtonsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetButtonWithResponse(ctx context.Context, buttonId string, reqEditors ...RequestEditorFn) (*GetButtonResponse, error) {
	args := c.Called(ctx, buttonId, reqEditors)
	return args.Get(0).(*GetButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateButtonWithBodyWithResponse(ctx context.Context, buttonId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateButtonResponse, error) {
	args := c.Called(ctx, buttonId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateButtonWithResponse(ctx context.Context, buttonId string, body UpdateButtonJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateButtonResponse, error) {
	args := c.Called(ctx, buttonId, body, reqEditors)
	return args.Get(0).(*UpdateButtonResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetCameraMotionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetCameraMotionsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetCameraMotionsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetCameraMotionWithResponse(ctx context.Context, cameraMotionId string, reqEditors ...RequestEditorFn) (*GetCameraMotionResponse, error) {
	args := c.Called(ctx, cameraMotionId, reqEditors)
	return args.Get(0).(*GetCameraMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateCameraMotionWithBodyWithResponse(ctx context.Context, cameraMotionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateCameraMotionResponse, error) {
	args := c.Called(ctx, cameraMotionId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateCameraMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateCameraMotionWithResponse(ctx context.Context, cameraMotionId string, body UpdateCameraMotionJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateCameraMotionResponse, error) {
	args := c.Called(ctx, cameraMotionId, body, reqEditors)
	return args.Get(0).(*UpdateCameraMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetContactsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetContactsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetContactsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetContactWithResponse(ctx context.Context, contactId string, reqEditors ...RequestEditorFn) (*GetContactResponse, error) {
	args := c.Called(ctx, contactId, reqEditors)
	return args.Get(0).(*GetContactResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateContactWithBodyWithResponse(ctx context.Context, contactId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateContactResponse, error) {
	args := c.Called(ctx, contactId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateContactResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateContactWithResponse(ctx context.Context, contactId string, body UpdateContactJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateContactResponse, error) {
	args := c.Called(ctx, contactId, body, reqEditors)
	return args.Get(0).(*UpdateContactResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetConvenienceAreaMotionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetConvenienceAreaMotionsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetConvenienceAreaMotionsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetConvenienceAreaMotionWithResponse(ctx context.Context, convenienceAreaMotionId string, reqEditors ...RequestEditorFn) (*GetConvenienceAreaMotionResponse, error) {
	args := c.Called(ctx, convenienceAreaMotionId, reqEditors)
	return args.Get(0).(*GetConvenienceAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateConvenienceAreaMotionWithBodyWithResponse(ctx context.Context, convenienceAreaMotionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateConvenienceAreaMotionResponse, error) {
	args := c.Called(ctx, convenienceAreaMotionId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateConvenienceAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateConvenienceAreaMotionWithResponse(ctx context.Context, convenienceAreaMotionId string, body UpdateConvenienceAreaMotionJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateConvenienceAreaMotionResponse, error) {
	args := c.Called(ctx, convenienceAreaMotionId, body, reqEditors)
	return args.Get(0).(*UpdateConvenienceAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDevicesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDevicesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetDevicesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) DeleteDeviceWithResponse(ctx context.Context, deviceId string, reqEditors ...RequestEditorFn) (*DeleteDeviceResponse, error) {
	args := c.Called(ctx, deviceId, reqEditors)
	return args.Get(0).(*DeleteDeviceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDeviceWithResponse(ctx context.Context, deviceId string, reqEditors ...RequestEditorFn) (*GetDeviceResponse, error) {
	args := c.Called(ctx, deviceId, reqEditors)
	return args.Get(0).(*GetDeviceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateDeviceWithBodyWithResponse(ctx context.Context, deviceId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateDeviceResponse, error) {
	args := c.Called(ctx, deviceId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateDeviceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateDeviceWithResponse(ctx context.Context, deviceId string, body UpdateDeviceJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateDeviceResponse, error) {
	args := c.Called(ctx, deviceId, body, reqEditors)
	return args.Get(0).(*UpdateDeviceResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDevicePowersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDevicePowersResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetDevicePowersResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDevicePowerWithResponse(ctx context.Context, deviceId string, reqEditors ...RequestEditorFn) (*GetDevicePowerResponse, error) {
	args := c.Called(ctx, deviceId, reqEditors)
	return args.Get(0).(*GetDevicePowerResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDeviceSoftwareUpdatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetDeviceSoftwareUpdatesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetDeviceSoftwareUpdatesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetDeviceSoftwareUpdateWithResponse(ctx context.Context, deviceSoftwareUpdateId string, reqEditors ...RequestEditorFn) (*GetDeviceSoftwareUpdateResponse, error) {
	args := c.Called(ctx, deviceSoftwareUpdateId, reqEditors)
	return args.Get(0).(*GetDeviceSoftwareUpdateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateDeviceSoftwareUpdateWithBodyWithResponse(ctx context.Context, deviceSoftwareUpdateId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateDeviceSoftwareUpdateResponse, error) {
	args := c.Called(ctx, deviceSoftwareUpdateId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateDeviceSoftwareUpdateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateDeviceSoftwareUpdateWithResponse(ctx context.Context, deviceSoftwareUpdateId string, body UpdateDeviceSoftwareUpdateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateDeviceSoftwareUpdateResponse, error) {
	args := c.Called(ctx, deviceSoftwareUpdateId, body, reqEditors)
	return args.Get(0).(*UpdateDeviceSoftwareUpdateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetEntertainmentsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetEntertainmentsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetEntertainmentsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetEntertainmentWithResponse(ctx context.Context, entertainmentId string, reqEditors ...RequestEditorFn) (*GetEntertainmentResponse, error) {
	args := c.Called(ctx, entertainmentId, reqEditors)
	return args.Get(0).(*GetEntertainmentResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateEntertainmentWithBodyWithResponse(ctx context.Context, entertainmentId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateEntertainmentResponse, error) {
	args := c.Called(ctx, entertainmentId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateEntertainmentResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateEntertainmentWithResponse(ctx context.Context, entertainmentId string, body UpdateEntertainmentJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateEntertainmentResponse, error) {
	args := c.Called(ctx, entertainmentId, body, reqEditors)
	return args.Get(0).(*UpdateEntertainmentResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetEntertainmentConfigurationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetEntertainmentConfigurationsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetEntertainmentConfigurationsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetEntertainmentConfigurationWithResponse(ctx context.Context, entertainmentConfigurationId string, reqEditors ...RequestEditorFn) (*GetEntertainmentConfigurationResponse, error) {
	args := c.Called(ctx, entertainmentConfigurationId, reqEditors)
	return args.Get(0).(*GetEntertainmentConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateEntertainmentConfigurationWithBodyWithResponse(ctx context.Context, entertainmentConfigurationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateEntertainmentConfigurationResponse, error) {
	args := c.Called(ctx, entertainmentConfigurationId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateEntertainmentConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateEntertainmentConfigurationWithResponse(ctx context.Context, entertainmentConfigurationId string, body UpdateEntertainmentConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateEntertainmentConfigurationResponse, error) {
	args := c.Called(ctx, entertainmentConfigurationId, body, reqEditors)
	return args.Get(0).(*UpdateEntertainmentConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGeofenceClientsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetGeofenceClientsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetGeofenceClientsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGeofenceClientWithResponse(ctx context.Context, geofenceClientId string, reqEditors ...RequestEditorFn) (*GetGeofenceClientResponse, error) {
	args := c.Called(ctx, geofenceClientId, reqEditors)
	return args.Get(0).(*GetGeofenceClientResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGeofenceClientWithBodyWithResponse(ctx context.Context, geofenceClientId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateGeofenceClientResponse, error) {
	args := c.Called(ctx, geofenceClientId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateGeofenceClientResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGeofenceClientWithResponse(ctx context.Context, geofenceClientId string, body UpdateGeofenceClientJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateGeofenceClientResponse, error) {
	args := c.Called(ctx, geofenceClientId, body, reqEditors)
	return args.Get(0).(*UpdateGeofenceClientResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGeolocationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetGeolocationsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetGeolocationsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGeolocationWithResponse(ctx context.Context, geolocationId string, reqEditors ...RequestEditorFn) (*GetGeolocationResponse, error) {
	args := c.Called(ctx, geolocationId, reqEditors)
	return args.Get(0).(*GetGeolocationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGeolocationWithBodyWithResponse(ctx context.Context, geolocationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateGeolocationResponse, error) {
	args := c.Called(ctx, geolocationId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateGeolocationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGeolocationWithResponse(ctx context.Context, geolocationId string, body UpdateGeolocationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateGeolocationResponse, error) {
	args := c.Called(ctx, geolocationId, body, reqEditors)
	return args.Get(0).(*UpdateGeolocationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedLightsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetGroupedLightsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetGroupedLightsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedLightWithResponse(ctx context.Context, groupedLightId string, reqEditors ...RequestEditorFn) (*GetGroupedLightResponse, error) {
	args := c.Called(ctx, groupedLightId, reqEditors)
	return args.Get(0).(*GetGroupedLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedLightWithBodyWithResponse(ctx context.Context, groupedLightId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateGroupedLightResponse, error) {
	args := c.Called(ctx, groupedLightId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateGroupedLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedLightWithResponse(ctx context.Context, groupedLightId string, body UpdateGroupedLightJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateGroupedLightResponse, error) {
	args := c.Called(ctx, groupedLightId, body, reqEditors)
	return args.Get(0).(*UpdateGroupedLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedLightLevelsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetGroupedLightLevelsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetGroupedLightLevelsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedLightLevelWithResponse(ctx context.Context, groupedLightLevelId string, reqEditors ...RequestEditorFn) (*GetGroupedLightLevelResponse, error) {
	args := c.Called(ctx, groupedLightLevelId, reqEditors)
	return args.Get(0).(*GetGroupedLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedLightLevelWithBodyWithResponse(ctx context.Context, groupedLightLevelId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateGroupedLightLevelResponse, error) {
	args := c.Called(ctx, groupedLightLevelId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateGroupedLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedLightLevelWithResponse(ctx context.Context, groupedLightLevelId string, body UpdateGroupedLightLevelJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateGroupedLightLevelResponse, error) {
	args := c.Called(ctx, groupedLightLevelId, body, reqEditors)
	return args.Get(0).(*UpdateGroupedLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedMotionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetGroupedMotionsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetGroupedMotionsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetGroupedMotionWithResponse(ctx context.Context, groupedMotionId string, reqEditors ...RequestEditorFn) (*GetGroupedMotionResponse, error) {
	args := c.Called(ctx, groupedMotionId, reqEditors)
	return args.Get(0).(*GetGroupedMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedMotionWithBodyWithResponse(ctx context.Context, groupedMotionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateGroupedMotionResponse, error) {
	args := c.Called(ctx, groupedMotionId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateGroupedMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateGroupedMotionWithResponse(ctx context.Context, groupedMotionId string, body UpdateGroupedMotionJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateGroupedMotionResponse, error) {
	args := c.Called(ctx, groupedMotionId, body, reqEditors)
	return args.Get(0).(*UpdateGroupedMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetHomekitsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHomekitsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetHomekitsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetHomekitWithResponse(ctx context.Context, homekitId string, reqEditors ...RequestEditorFn) (*GetHomekitResponse, error) {
	args := c.Called(ctx, homekitId, reqEditors)
	return args.Get(0).(*GetHomekitResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateHomekitWithBodyWithResponse(ctx context.Context, homekitId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateHomekitResponse, error) {
	args := c.Called(ctx, homekitId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateHomekitResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateHomekitWithResponse(ctx context.Context, homekitId string, body UpdateHomekitJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateHomekitResponse, error) {
	args := c.Called(ctx, homekitId, body, reqEditors)
	return args.Get(0).(*UpdateHomekitResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetLightsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetLightsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetLightsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetLightWithResponse(ctx context.Context, lightId string, reqEditors ...RequestEditorFn) (*GetLightResponse, error) {
	args := c.Called(ctx, lightId, reqEditors)
	return args.Get(0).(*GetLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateLightWithBodyWithResponse(ctx context.Context, lightId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLightResponse, error) {
	args := c.Called(ctx, lightId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateLightWithResponse(ctx context.Context, lightId string, body UpdateLightJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLightResponse, error) {
	args := c.Called(ctx, lightId, body, reqEditors)
	return args.Get(0).(*UpdateLightResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetLightLevelsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetLightLevelsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetLightLevelsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetLightLevelWithResponse(ctx context.Context, lightId string, reqEditors ...RequestEditorFn) (*GetLightLevelResponse, error) {
	args := c.Called(ctx, lightId, reqEditors)
	return args.Get(0).(*GetLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateLightLevelWithBodyWithResponse(ctx context.Context, lightId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateLightLevelResponse, error) {
	args := c.Called(ctx, lightId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateLightLevelWithResponse(ctx context.Context, lightId string, body UpdateLightLevelJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateLightLevelResponse, error) {
	args := c.Called(ctx, lightId, body, reqEditors)
	return args.Get(0).(*UpdateLightLevelResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMattersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMattersResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetMattersResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMatterWithResponse(ctx context.Context, matterId string, reqEditors ...RequestEditorFn) (*GetMatterResponse, error) {
	args := c.Called(ctx, matterId, reqEditors)
	return args.Get(0).(*GetMatterResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMatterWithBodyWithResponse(ctx context.Context, matterId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateMatterResponse, error) {
	args := c.Called(ctx, matterId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateMatterResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMatterWithResponse(ctx context.Context, matterId string, body UpdateMatterJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateMatterResponse, error) {
	args := c.Called(ctx, matterId, body, reqEditors)
	return args.Get(0).(*UpdateMatterResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMatterFabricsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMatterFabricsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetMatterFabricsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMatterFabricWithResponse(ctx context.Context, matterFabricId string, reqEditors ...RequestEditorFn) (*GetMatterFabricResponse, error) {
	args := c.Called(ctx, matterFabricId, reqEditors)
	return args.Get(0).(*GetMatterFabricResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMatterFabricWithBodyWithResponse(ctx context.Context, matterFabricId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateMatterFabricResponse, error) {
	args := c.Called(ctx, matterFabricId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateMatterFabricResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMatterFabricWithResponse(ctx context.Context, matterFabricId string, body UpdateMatterFabricJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateMatterFabricResponse, error) {
	args := c.Called(ctx, matterFabricId, body, reqEditors)
	return args.Get(0).(*UpdateMatterFabricResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionSensorsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMotionSensorsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetMotionSensorsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionSensorWithResponse(ctx context.Context, motionId string, reqEditors ...RequestEditorFn) (*GetMotionSensorResponse, error) {
	args := c.Called(ctx, motionId, reqEditors)
	return args.Get(0).(*GetMotionSensorResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionSensorWithBodyWithResponse(ctx context.Context, motionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateMotionSensorResponse, error) {
	args := c.Called(ctx, motionId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateMotionSensorResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionSensorWithResponse(ctx context.Context, motionId string, body UpdateMotionSensorJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateMotionSensorResponse, error) {
	args := c.Called(ctx, motionId, body, reqEditors)
	return args.Get(0).(*UpdateMotionSensorResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionAreaCandidatesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMotionAreaCandidatesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetMotionAreaCandidatesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionAreaCandidateWithResponse(ctx context.Context, motionAreaCandidateId string, reqEditors ...RequestEditorFn) (*GetMotionAreaCandidateResponse, error) {
	args := c.Called(ctx, motionAreaCandidateId, reqEditors)
	return args.Get(0).(*GetMotionAreaCandidateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionAreaCandidateWithBodyWithResponse(ctx context.Context, motionAreaCandidateId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateMotionAreaCandidateResponse, error) {
	args := c.Called(ctx, motionAreaCandidateId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateMotionAreaCandidateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionAreaCandidateWithResponse(ctx context.Context, motionAreaCandidateId string, body UpdateMotionAreaCandidateJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateMotionAreaCandidateResponse, error) {
	args := c.Called(ctx, motionAreaCandidateId, body, reqEditors)
	return args.Get(0).(*UpdateMotionAreaCandidateResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionAreaConfigurationsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetMotionAreaConfigurationsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetMotionAreaConfigurationsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetMotionAreaConfigurationWithResponse(ctx context.Context, motionAreaConfigurationId string, reqEditors ...RequestEditorFn) (*GetMotionAreaConfigurationResponse, error) {
	args := c.Called(ctx, motionAreaConfigurationId, reqEditors)
	return args.Get(0).(*GetMotionAreaConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionAreaConfigurationWithBodyWithResponse(ctx context.Context, motionAreaConfigurationId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateMotionAreaConfigurationResponse, error) {
	args := c.Called(ctx, motionAreaConfigurationId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateMotionAreaConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateMotionAreaConfigurationWithResponse(ctx context.Context, motionAreaConfigurationId string, body UpdateMotionAreaConfigurationJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateMotionAreaConfigurationResponse, error) {
	args := c.Called(ctx, motionAreaConfigurationId, body, reqEditors)
	return args.Get(0).(*UpdateMotionAreaConfigurationResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetRelativeRotariesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRelativeRotariesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetRelativeRotariesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetRelativeRotaryWithResponse(ctx context.Context, relativeRotaryId string, reqEditors ...RequestEditorFn) (*GetRelativeRotaryResponse, error) {
	args := c.Called(ctx, relativeRotaryId, reqEditors)
	return args.Get(0).(*GetRelativeRotaryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateRelativeRotaryWithBodyWithResponse(ctx context.Context, relativeRotaryId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateRelativeRotaryResponse, error) {
	args := c.Called(ctx, relativeRotaryId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateRelativeRotaryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateRelativeRotaryWithResponse(ctx context.Context, relativeRotaryId string, body UpdateRelativeRotaryJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateRelativeRotaryResponse, error) {
	args := c.Called(ctx, relativeRotaryId, body, reqEditors)
	return args.Get(0).(*UpdateRelativeRotaryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetRoomsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRoomsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetRoomsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateRoomWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateRoomResponse, error) {
	args := c.Called(ctx, contentType, body, reqEditors)
	return args.Get(0).(*CreateRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateRoomWithResponse(ctx context.Context, body CreateRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateRoomResponse, error) {
	args := c.Called(ctx, body, reqEditors)
	return args.Get(0).(*CreateRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) DeleteRoomWithResponse(ctx context.Context, roomId string, reqEditors ...RequestEditorFn) (*DeleteRoomResponse, error) {
	args := c.Called(ctx, roomId, reqEditors)
	return args.Get(0).(*DeleteRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetRoomWithResponse(ctx context.Context, roomId string, reqEditors ...RequestEditorFn) (*GetRoomResponse, error) {
	args := c.Called(ctx, roomId, reqEditors)
	return args.Get(0).(*GetRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateRoomWithBodyWithResponse(ctx context.Context, roomId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateRoomResponse, error) {
	args := c.Called(ctx, roomId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateRoomWithResponse(ctx context.Context, roomId string, body UpdateRoomJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateRoomResponse, error) {
	args := c.Called(ctx, roomId, body, reqEditors)
	return args.Get(0).(*UpdateRoomResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetScenesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetScenesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetScenesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateSceneWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSceneResponse, error) {
	args := c.Called(ctx, contentType, body, reqEditors)
	return args.Get(0).(*CreateSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateSceneWithResponse(ctx context.Context, body CreateSceneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSceneResponse, error) {
	args := c.Called(ctx, body, reqEditors)
	return args.Get(0).(*CreateSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) DeleteSceneWithResponse(ctx context.Context, sceneId string, reqEditors ...RequestEditorFn) (*DeleteSceneResponse, error) {
	args := c.Called(ctx, sceneId, reqEditors)
	return args.Get(0).(*DeleteSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSceneWithResponse(ctx context.Context, sceneId string, reqEditors ...RequestEditorFn) (*GetSceneResponse, error) {
	args := c.Called(ctx, sceneId, reqEditors)
	return args.Get(0).(*GetSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSceneWithBodyWithResponse(ctx context.Context, sceneId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSceneResponse, error) {
	args := c.Called(ctx, sceneId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSceneWithResponse(ctx context.Context, sceneId string, body UpdateSceneJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSceneResponse, error) {
	args := c.Called(ctx, sceneId, body, reqEditors)
	return args.Get(0).(*UpdateSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSecurityAreaMotionsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSecurityAreaMotionsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetSecurityAreaMotionsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSecurityAreaMotionWithResponse(ctx context.Context, securityAreaMotionId string, reqEditors ...RequestEditorFn) (*GetSecurityAreaMotionResponse, error) {
	args := c.Called(ctx, securityAreaMotionId, reqEditors)
	return args.Get(0).(*GetSecurityAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSecurityAreaMotionWithBodyWithResponse(ctx context.Context, securityAreaMotionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSecurityAreaMotionResponse, error) {
	args := c.Called(ctx, securityAreaMotionId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateSecurityAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSecurityAreaMotionWithResponse(ctx context.Context, securityAreaMotionId string, body UpdateSecurityAreaMotionJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSecurityAreaMotionResponse, error) {
	args := c.Called(ctx, securityAreaMotionId, body, reqEditors)
	return args.Get(0).(*UpdateSecurityAreaMotionResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetServiceGroupsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetServiceGroupsResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetServiceGroupsResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetServiceGroupWithResponse(ctx context.Context, serviceGroupId string, reqEditors ...RequestEditorFn) (*GetServiceGroupResponse, error) {
	args := c.Called(ctx, serviceGroupId, reqEditors)
	return args.Get(0).(*GetServiceGroupResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateServiceGroupWithBodyWithResponse(ctx context.Context, serviceGroupId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateServiceGroupResponse, error) {
	args := c.Called(ctx, serviceGroupId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateServiceGroupResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateServiceGroupWithResponse(ctx context.Context, serviceGroupId string, body UpdateServiceGroupJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateServiceGroupResponse, error) {
	args := c.Called(ctx, serviceGroupId, body, reqEditors)
	return args.Get(0).(*UpdateServiceGroupResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSmartScenesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSmartScenesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetSmartScenesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateSmartSceneWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSmartSceneResponse, error) {
	args := c.Called(ctx, contentType, body, reqEditors)
	return args.Get(0).(*CreateSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateSmartSceneWithResponse(ctx context.Context, body CreateSmartSceneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSmartSceneResponse, error) {
	args := c.Called(ctx, body, reqEditors)
	return args.Get(0).(*CreateSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) DeleteSmartSceneWithResponse(ctx context.Context, sceneId string, reqEditors ...RequestEditorFn) (*DeleteSmartSceneResponse, error) {
	args := c.Called(ctx, sceneId, reqEditors)
	return args.Get(0).(*DeleteSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSmartSceneWithResponse(ctx context.Context, sceneId string, reqEditors ...RequestEditorFn) (*GetSmartSceneResponse, error) {
	args := c.Called(ctx, sceneId, reqEditors)
	return args.Get(0).(*GetSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSmartSceneWithBodyWithResponse(ctx context.Context, sceneId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSmartSceneResponse, error) {
	args := c.Called(ctx, sceneId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSmartSceneWithResponse(ctx context.Context, sceneId string, body UpdateSmartSceneJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSmartSceneResponse, error) {
	args := c.Called(ctx, sceneId, body, reqEditors)
	return args.Get(0).(*UpdateSmartSceneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSpeakersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetSpeakersResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetSpeakersResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetSpeakerWithResponse(ctx context.Context, speakerId string, reqEditors ...RequestEditorFn) (*GetSpeakerResponse, error) {
	args := c.Called(ctx, speakerId, reqEditors)
	return args.Get(0).(*GetSpeakerResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSpeakerWithBodyWithResponse(ctx context.Context, speakerId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateSpeakerResponse, error) {
	args := c.Called(ctx, speakerId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateSpeakerResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateSpeakerWithResponse(ctx context.Context, speakerId string, body UpdateSpeakerJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateSpeakerResponse, error) {
	args := c.Called(ctx, speakerId, body, reqEditors)
	return args.Get(0).(*UpdateSpeakerResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetTampersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTampersResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetTampersResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetTamperWithResponse(ctx context.Context, tamperId string, reqEditors ...RequestEditorFn) (*GetTamperResponse, error) {
	args := c.Called(ctx, tamperId, reqEditors)
	return args.Get(0).(*GetTamperResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateTamperWithBodyWithResponse(ctx context.Context, tamperId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTamperResponse, error) {
	args := c.Called(ctx, tamperId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateTamperResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateTamperWithResponse(ctx context.Context, tamperId string, body UpdateTamperJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTamperResponse, error) {
	args := c.Called(ctx, tamperId, body, reqEditors)
	return args.Get(0).(*UpdateTamperResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetTemperaturesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTemperaturesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetTemperaturesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetTemperatureWithResponse(ctx context.Context, temperatureId string, reqEditors ...RequestEditorFn) (*GetTemperatureResponse, error) {
	args := c.Called(ctx, temperatureId, reqEditors)
	return args.Get(0).(*GetTemperatureResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateTemperatureWithBodyWithResponse(ctx context.Context, temperatureId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateTemperatureResponse, error) {
	args := c.Called(ctx, temperatureId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateTemperatureResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateTemperatureWithResponse(ctx context.Context, temperatureId string, body UpdateTemperatureJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateTemperatureResponse, error) {
	args := c.Called(ctx, temperatureId, body, reqEditors)
	return args.Get(0).(*UpdateTemperatureResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetWifiConnectivitiesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetWifiConnectivitiesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetWifiConnectivitiesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetWifiConnectivityWithResponse(ctx context.Context, wifiConnectivityId string, reqEditors ...RequestEditorFn) (*GetWifiConnectivityResponse, error) {
	args := c.Called(ctx, wifiConnectivityId, reqEditors)
	return args.Get(0).(*GetWifiConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateWifiConnectivityWithBodyWithResponse(ctx context.Context, wifiConnectivityId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateWifiConnectivityResponse, error) {
	args := c.Called(ctx, wifiConnectivityId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateWifiConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateWifiConnectivityWithResponse(ctx context.Context, wifiConnectivityId string, body UpdateWifiConnectivityJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateWifiConnectivityResponse, error) {
	args := c.Called(ctx, wifiConnectivityId, body, reqEditors)
	return args.Get(0).(*UpdateWifiConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZgpConnectivitiesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetZgpConnectivitiesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetZgpConnectivitiesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZgpConnectivityWithResponse(ctx context.Context, zgpConnectivityId string, reqEditors ...RequestEditorFn) (*GetZgpConnectivityResponse, error) {
	args := c.Called(ctx, zgpConnectivityId, reqEditors)
	return args.Get(0).(*GetZgpConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZgpConnectivityWithBodyWithResponse(ctx context.Context, zgpConnectivityId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateZgpConnectivityResponse, error) {
	args := c.Called(ctx, zgpConnectivityId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateZgpConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZgpConnectivityWithResponse(ctx context.Context, zgpConnectivityId string, body UpdateZgpConnectivityJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateZgpConnectivityResponse, error) {
	args := c.Called(ctx, zgpConnectivityId, body, reqEditors)
	return args.Get(0).(*UpdateZgpConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZigbeeConnectivitiesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetZigbeeConnectivitiesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetZigbeeConnectivitiesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZigbeeConnectivityWithResponse(ctx context.Context, zigbeeConnectivityId string, reqEditors ...RequestEditorFn) (*GetZigbeeConnectivityResponse, error) {
	args := c.Called(ctx, zigbeeConnectivityId, reqEditors)
	return args.Get(0).(*GetZigbeeConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZigbeeConnectivityWithBodyWithResponse(ctx context.Context, zigbeeConnectivityId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateZigbeeConnectivityResponse, error) {
	args := c.Called(ctx, zigbeeConnectivityId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateZigbeeConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZigbeeConnectivityWithResponse(ctx context.Context, zigbeeConnectivityId string, body UpdateZigbeeConnectivityJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateZigbeeConnectivityResponse, error) {
	args := c.Called(ctx, zigbeeConnectivityId, body, reqEditors)
	return args.Get(0).(*UpdateZigbeeConnectivityResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZigbeeDeviceDiscoveriesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetZigbeeDeviceDiscoveriesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetZigbeeDeviceDiscoveriesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZigbeeDeviceDiscoveryWithResponse(ctx context.Context, zigbeeDeviceDiscoveryId string, reqEditors ...RequestEditorFn) (*GetZigbeeDeviceDiscoveryResponse, error) {
	args := c.Called(ctx, zigbeeDeviceDiscoveryId, reqEditors)
	return args.Get(0).(*GetZigbeeDeviceDiscoveryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZigbeeDeviceDiscoveryWithBodyWithResponse(ctx context.Context, zigbeeDeviceDiscoveryId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateZigbeeDeviceDiscoveryResponse, error) {
	args := c.Called(ctx, zigbeeDeviceDiscoveryId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateZigbeeDeviceDiscoveryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZigbeeDeviceDiscoveryWithResponse(ctx context.Context, zigbeeDeviceDiscoveryId string, body UpdateZigbeeDeviceDiscoveryJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateZigbeeDeviceDiscoveryResponse, error) {
	args := c.Called(ctx, zigbeeDeviceDiscoveryId, body, reqEditors)
	return args.Get(0).(*UpdateZigbeeDeviceDiscoveryResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZonesWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetZonesResponse, error) {
	args := c.Called(ctx, reqEditors)
	return args.Get(0).(*GetZonesResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateZoneWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateZoneResponse, error) {
	args := c.Called(ctx, contentType, body, reqEditors)
	return args.Get(0).(*CreateZoneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) CreateZoneWithResponse(ctx context.Context, body CreateZoneJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateZoneResponse, error) {
	args := c.Called(ctx, body, reqEditors)
	return args.Get(0).(*CreateZoneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) DeleteZoneWithResponse(ctx context.Context, zoneId string, reqEditors ...RequestEditorFn) (*DeleteZoneResponse, error) {
	args := c.Called(ctx, zoneId, reqEditors)
	return args.Get(0).(*DeleteZoneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) GetZoneWithResponse(ctx context.Context, zoneId string, reqEditors ...RequestEditorFn) (*GetZoneResponse, error) {
	args := c.Called(ctx, zoneId, reqEditors)
	return args.Get(0).(*GetZoneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZoneWithBodyWithResponse(ctx context.Context, zoneId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateZoneResponse, error) {
	args := c.Called(ctx, zoneId, contentType, body, reqEditors)
	return args.Get(0).(*UpdateZoneResponse), args.Error(1)
}

func (c *ClientWithResponsesMock) UpdateZoneWithResponse(ctx context.Context, zoneId string, body UpdateZoneJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateZoneResponse, error) {
	args := c.Called(ctx, zoneId, body, reqEditors)
	return args.Get(0).(*UpdateZoneResponse), args.Error(1)
}

