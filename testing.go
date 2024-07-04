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
	args := c.Called(ctx, bridgeId, reqEditors)
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
