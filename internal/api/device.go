package api

import (
	"github.com/SBAustin951/aurorasync/internal/models"
	"github.com/google/uuid"
)

const (
	allDevicesEndpoint  = "/router/api/v1/user/devices"
	deviceStateEndpoint = "/router/api/v1/device/state"
	controlEndpoint     = "/router/api/v1/device/control"
)

func (g *GoveeClient) GetAllDevices() (*models.DeviceResponse, error) {
	var deviceResponse models.DeviceResponse
	if err := g.Get(allDevicesEndpoint, deviceResponse); err != nil {
		return nil, err
	}
	return &deviceResponse, nil
}

func (g *GoveeClient) GetDeviceState(sku, deviceId string) (*models.DeviceStateResponse, error) {
	deviceStateRequest := models.DeviceStateRequest{
		RequestID: uuid.New(),
		Payload: models.StatePayload{
			SKU:      sku,
			DeviceID: deviceId,
		},
	}
	var deviceStateResponse models.DeviceStateResponse
	if err := g.Post(deviceStateEndpoint, deviceStateRequest, deviceStateResponse); err != nil {
		return nil, err
	}
	return &deviceStateResponse, nil
}

func (g *GoveeClient) PowerOn(sku, deviceId string) (*models.ControlResponse, error) {
	value := new(int)
	*value = 1
	controlRequest := models.ControlRequest{
		RequestID: uuid.New(),
		Payload: models.ControlPayload{
			SKU:      sku,
			DeviceID: deviceId,
			Capability: models.Capability{
				Type:     "devices.capabilities.on_off",
				Instance: "powerSwitch",
				Value:    value,
			},
		},
	}
	var controlResponse models.ControlResponse
	if err := g.Post(controlEndpoint, controlRequest, controlResponse); err != nil {
		return nil, err
	}
	return &controlResponse, nil
}

func (g *GoveeClient) PowerOff(sku, deviceId string) (*models.ControlResponse, error) {
	value := new(int)
	*value = 1
	controlRequest := models.ControlRequest{
		RequestID: uuid.New(),
		Payload: models.ControlPayload{
			SKU:      sku,
			DeviceID: deviceId,
			Capability: models.Capability{
				Type:     "devices.capabilities.on_off",
				Instance: "powerSwitch",
				Value:    value,
			},
		},
	}
	var controlResponse models.ControlResponse
	if err := g.Post(controlEndpoint, controlRequest, controlResponse); err != nil {
		return nil, err
	}
	return &controlResponse, nil
}
