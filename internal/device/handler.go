package main

import (
	"context"
	device "github.com/OpenIoT-Hub/openiot-server/kitex_gen/device"
)

// DeviceServiceImpl implements the last service interface defined in the IDL.
type DeviceServiceImpl struct{}

// Ping implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) Ping(ctx context.Context, req *device.PingRequst) (resp *device.PingResponse, err error) {
	// TODO: Your code here...
	return
}

// Create implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) Create(ctx context.Context, req *device.CreateDeviceRequst) (resp *device.CreateDeviceResponse, err error) {
	// TODO: Your code here...
	return
}

// Remove implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) Remove(ctx context.Context, req *device.RemoveDeviceRequst) (resp *device.RemoveDeviceResponse, err error) {
	// TODO: Your code here...
	return
}

// Update implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) Update(ctx context.Context, req *device.UpdateDeviceRequst) (resp *device.UpdateDeviceResponse, err error) {
	// TODO: Your code here...
	return
}

// GetDeviceByID implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) GetDeviceByID(ctx context.Context, req *device.GetDeviceRequst) (resp *device.GetDeviceResponse, err error) {
	// TODO: Your code here...
	return
}

// ListDevice implements the DeviceServiceImpl interface.
func (s *DeviceServiceImpl) ListDevice(ctx context.Context, req *device.ListDeviceRequst) (resp *device.ListDeviceResponse, err error) {
	// TODO: Your code here...
	return
}
